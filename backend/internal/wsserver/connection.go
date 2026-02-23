package wsserver

// {"type":"auth","data":"{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VydGFnIjoiYWxpY2VfZGV2IiwidXNlcm5hbWUiOiJBbGljZSIsImlzcyI6InlvdXItYXBwLW5hbWUiLCJleHAiOjE3NzIxMTc3NDYsIm5iZiI6MTc3MTg1ODU0NiwiaWF0IjoxNzcxODU4NTQ2fQ.eLeHzE4B-KUB0Aq_C_fsKiIoU90aRsFzPM0eRM6sqXA\"}"}

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	jwttoken "backend/pkg/jwt_token"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Postgres interface {
	GetChatsIDByUserID(ctx context.Context, userID int) ([]int, error)
	CreateMessage(ctx context.Context, message domain.MessageCreate) (int, error)
}

type WsServer struct {
	postgres    Postgres
	jwtManager  *jwttoken.JWTManager
	upgrader    websocket.Upgrader
	broadcast   chan *BroadcastMessage
	connections map[*websocket.Conn]Client
	rooms       map[int]map[*websocket.Conn]struct{}
	mutex       *sync.RWMutex
}

func New(postgres *postgres.Pool, jwtManager *jwttoken.JWTManager) *WsServer {
	return &WsServer{
		postgres:   postgres,
		jwtManager: jwtManager,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				if r.Header.Get("Origin") == "http://localhost:5173" {
					return true
				}
				//return false
				return true
			},
		},
		broadcast:   make(chan *BroadcastMessage),
		connections: make(map[*websocket.Conn]Client),
		rooms:       make(map[int]map[*websocket.Conn]struct{}),
		mutex:       &sync.RWMutex{},
	}
}

func (ws *WsServer) Handler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Start connection ...")
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Error with websocket connection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ws.mutex.Lock()
	ws.connections[conn] = *NewDefaultClient(conn)
	ws.mutex.Unlock()

	go ws.readFromClient(conn)
}

func (ws *WsServer) Start() {
	go ws.runBroadcast()
}

func (ws *WsServer) readFromClient(conn *websocket.Conn) {
	logrus.Info("Start readFromClient")
	for {
		msg := new(Message)
		if err := conn.ReadJSON(msg); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Warnf("websocket unexpected close: %v", err)
			}
			logrus.Errorf("websocket unexpected close: %v 1", err)
			break
		}
		logrus.Info("Start process msg ...")
		switch msg.Type {
		case TypeAuth:
			var data AuthData
			if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
				logrus.Warnf("auth: invalid data: %v", err)
				//TODO: отправить ошибку
				break
			}
			if err := ws.authClient(conn, data); err != nil {
				logrus.Warnf("auth: %v", err)
				//TODO: отправить ошибку
			}
			logrus.Info("succses auth")
		case TypeChatMessage:
			ws.mutex.RLock()
			client := ws.connections[conn]
			ws.mutex.RUnlock()
			if client.UserID == 0 {
				//TODO: вернуть ошибку
				break
			}
			var data ChatMessageData
			if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
				logrus.Warnf("chat message: invalid data: %v", err)
				//TODO: отправить ошибку
				break
			}
			broadcastMessage := BroadcastMessage{
				UserID:   client.UserID,
				Username: client.Username,
				Usertag:  client.Usertag,
				ChatID:   data.ChatID,
				Text:     data.Text,
				Time:     time.Now(),
			}
			ws.broadcast <- &broadcastMessage
		}
	}
	ws.mutex.Lock()
	client_rooms := ws.connections[conn].Rooms
	delete(ws.connections, conn)
	for chat_id := range client_rooms {
		delete(ws.rooms[chat_id], conn)
	}
	ws.mutex.Unlock()
}

func (ws *WsServer) authClient(conn *websocket.Conn, data AuthData) error {
	logrus.Info("Start auth user ...")
	userClaims, err := ws.jwtManager.ParseToken(data.JWTToken)
	if err != nil {
		return fmt.Errorf("parse JWT token: %w", err)
	}
	chatIDs, err := ws.postgres.GetChatsIDByUserID(context.Background(), userClaims.UserID)
	if err != nil {
		return fmt.Errorf("get chat ids by user: %w", err)
	}

	logrus.Info("Parse token, get info in DB, start add connection")

	ws.mutex.Lock()
	for _, chatID := range chatIDs {
		if ws.rooms[chatID] == nil {
			ws.rooms[chatID] = make(map[*websocket.Conn]struct{})
		}
		ws.rooms[chatID][conn] = struct{}{}
	}
	client := ws.connections[conn]
	client.UserID = userClaims.UserID
	client.Username = userClaims.Username
	client.Usertag = userClaims.Usertag
	client.Rooms = chatIDs
	ws.connections[conn] = client
	ws.mutex.Unlock()

	return nil
}

func (ws *WsServer) runBroadcast() {
	for msg := range ws.broadcast {
		logrus.Info("Start process msg with broadcast ...")
		messageCreate := domain.NewMessageCreate(msg.UserID, msg.ChatID, msg.Text, msg.Time)
		msgID, err := ws.postgres.CreateMessage(context.Background(), messageCreate)
		if err != nil {
			logrus.Errorf("create message: %v", err)
			continue
		}
		msg.MessageID = msgID
		conns := ws.getRoomConnections(msg.ChatID)
		if conns == nil {
			continue
		}
		for _, conn := range conns {
			if err := conn.WriteJSON(msg); err != nil {
				logrus.Warnf("write broadcast message: %v", err)
			}
		}
	}
}

func (ws *WsServer) getRoomConnections(chatID int) []*websocket.Conn {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	room, ok := ws.rooms[chatID]
	if !ok || len(room) == 0 {
		logrus.Warnf("no connections for chat: %d", chatID)
		return nil
	}
	connList := make([]*websocket.Conn, 0, len(room))
	for conn := range room {
		connList = append(connList, conn)
	}
	return connList
}

func (ws *WsServer) Close() {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	for conn := range ws.connections {
		_ = conn.WriteControl(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseGoingAway, "server shutdown"),
			time.Now().Add(5*time.Second),
		)
		_ = conn.Close()
	}

	ws.connections = make(map[*websocket.Conn]Client)
	ws.rooms = make(map[int]map[*websocket.Conn]struct{})

	close(ws.broadcast)
}
