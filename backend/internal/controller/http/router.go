package wsserver

import (
	"backend/internal/middleware"
	"backend/internal/user/register_user"
	"backend/internal/wsserver"
	"net/http"
)

func Router(ws *wsserver.WsServer) http.Handler {

	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()
	apiV1Mux := http.NewServeMux()
	privateMux := http.NewServeMux()

	mainMux.HandleFunc("/ws", ws.WsHandler)

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	apiMux.Handle("/v1/", http.StripPrefix("/v1", apiV1Mux))

	privateMux.HandleFunc("/registration", register_user.HTTP_V1)

	privateHandler := middleware.AuthMiddleware()(privateMux)
	apiV1Mux.Handle("/private/", http.StripPrefix("/private", privateHandler))

	handler := middleware.CORSMiddleware([]string{"http://localhost:5173"})(mainMux)

	return handler
}
