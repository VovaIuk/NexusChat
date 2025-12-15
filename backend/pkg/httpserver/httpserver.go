package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Config struct {
	Port string `default:"8080" envconfig:"HTTP_PORT"`
}

type Server struct {
	server *http.Server
}

func New(handler http.Handler, c Config) *Server {
	server := &http.Server{
		Addr:         ":" + c.Port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{server: server}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Close корректно завершает работу сервера
func (s *Server) Close(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
