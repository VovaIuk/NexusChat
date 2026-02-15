package wsserver

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsServer struct {
	upgrader  *websocket.Upgrader
	rooms     map[int][]int // id чата: список пользователей
	clients   map[*websocket.Conn]struct{}
	mutex     *sync.RWMutex
	broadcast chan *wsMessage
}

func NewWSServers() *WsServer {
	return &WsServer{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // потом сделать
			},
		},
		clients:   map[*websocket.Conn]struct{}{},
		rooms:     map[int][]int{},
		mutex:     &sync.RWMutex{},
		broadcast: make(chan *wsMessage),
	}
}

func (ws *WsServer) Start() {
	go ws.writeToClientsBroadcast()
}

func (ws *WsServer) Close() {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	for conn := range ws.clients {
		if err := conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseGoingAway, "server shutdown"),
			time.Now().Add(5*time.Second)); err != nil {
			logrus.Warnf("Error sending close message: %v", err)
		}
		if err := conn.Close(); err != nil {
			logrus.Errorf("Error closing connection: %v", err)
		}
	}

	ws.clients = make(map[*websocket.Conn]struct{})

	defer func() {
		if recover() != nil {
			logrus.Warn("Broadcast channel already closed")
		}
	}()
	close(ws.broadcast)
}

func (ws *WsServer) WsHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("web socket conn...")
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Error with websocket connection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logrus.Info(conn.RemoteAddr())
	ws.mutex.Lock()

	ws.clients[conn] = struct{}{}
	ws.mutex.Unlock()
	go ws.readFromClient(conn)
}

func (ws *WsServer) readFromClient(conn *websocket.Conn) {
	for {
		msg := new(wsMessage)
		if err := conn.ReadJSON(msg); err != nil {
			logrus.Errorf("Error with reading WebSocket: %v", err)
			break
		}
		logrus.Info(msg)
		msg.Time = time.Now().Format("15.04")
		logrus.Debug(msg)
		ws.broadcast <- msg
	}
	ws.mutex.Lock()
	delete(ws.clients, conn)
	ws.mutex.Unlock()
}

func (ws *WsServer) writeToClientsBroadcast() {
	for msg := range ws.broadcast {
		ws.mutex.RLock()
		for conn := range ws.clients {
			if err := conn.WriteJSON(msg); err != nil {
				logrus.Errorf("Error with writing message: %v", err)
			}
		}
		ws.mutex.RUnlock()
	}
}
