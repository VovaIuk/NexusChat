package wsserver

import (
	_ "embed"
	"net/http"
	"sync"
	"time"

	"backend/internal/middleware"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type wsServer struct {
	mux       *http.ServeMux
	server    *http.Server
	upgrader  *websocket.Upgrader
	clients   map[*websocket.Conn]struct{}
	mutex     *sync.RWMutex
	broadcast chan *wsMessage
}

type WSServer interface {
	Start() error
}

func NewWsServer(addr string) WSServer {
	mux := http.NewServeMux()
	handler := middleware.CORSMiddleware([]string{"http://localhost:5173"})(mux)
	return &wsServer{
		mux: mux,
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // потом сделать
			},
		},
		clients:   map[*websocket.Conn]struct{}{},
		mutex:     &sync.RWMutex{},
		broadcast: make(chan *wsMessage),
	}
}

func (ws *wsServer) Start() error {
	ws.mux.HandleFunc("/api/v1/test", ws.HandlerTest)
	ws.mux.HandleFunc("/ws", ws.wsHandler)
	ws.mux.Handle("/docs/", http.StripPrefix("/docs/",
		http.FileServer(http.Dir("internal/docs")),
	))
	go ws.writeToClientsBroadcast()
	return ws.server.ListenAndServe()
}

func (ws *wsServer) HandlerTest(w http.ResponseWriter, r *http.Request) {
	logrus.Info("test handler")
	w.Write([]byte("Test handler"))
}

func (ws *wsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
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

func (ws *wsServer) readFromClient(conn *websocket.Conn) {
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

func (ws *wsServer) writeToClientsBroadcast() {
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
