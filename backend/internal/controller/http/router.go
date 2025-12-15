package httpcontroller

import (
	"backend/internal/middleware"
	"backend/internal/user/login_user"
	"backend/internal/user/register_user"
	"backend/internal/wsserver"
	"embed"
	"io/fs"
	"net/http"

	"github.com/sirupsen/logrus"
)

//go:embed docs
var docsFS embed.FS

func Router(ws *wsserver.WsServer) http.Handler {

	sub, err := fs.Sub(docsFS, "docs")
	if err != nil {
		logrus.Fatal("Failed to create sub FS for docs:", err)
	}
	docsSubFS := http.FS(sub)

	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()
	apiV1Mux := http.NewServeMux()
	privateMux := http.NewServeMux()

	mainMux.HandleFunc("/ws", ws.WsHandler)

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	apiMux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(docsSubFS)))

	apiMux.Handle("/v1/", http.StripPrefix("/v1", apiV1Mux))

	apiV1Mux.HandleFunc("/login", login_user.HTTP_V1)
	apiV1Mux.HandleFunc("/registration", register_user.HTTP_V1)

	privateHandler := middleware.AuthMiddleware()(privateMux)
	apiV1Mux.Handle("/private/", http.StripPrefix("/private", privateHandler))

	handler := middleware.CORSMiddleware([]string{"http://localhost:5173", "http://localhost:8004"})(mainMux)

	return handler
}
