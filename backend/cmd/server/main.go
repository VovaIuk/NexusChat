package main

import (
	"backend/config"
	"backend/internal/adapter/postgres"
	httpcontroller "backend/internal/controller/http"
	"backend/internal/middleware"
	"backend/internal/user/login_user"
	"backend/internal/user/register_user"
	"backend/internal/wsserver"
	"backend/pkg/httpserver"
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	c, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	err = AppRun(context.Background(), c)
	if err != nil {
		panic(err)
	}
}

func AppRun(ctx context.Context, c config.Config) error {
	pgPool, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	jwtManager := middleware.InitAuth(c.JWT)

	wsServer := wsserver.NewWSServers()
	wsServer.Start()

	login_user.New(pgPool, jwtManager)
	register_user.New(pgPool)

	router := httpcontroller.Router(wsServer)
	server := httpserver.New(router, c.HTTP)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("Server failed: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Close(shutdownCtx); err != nil {
		logrus.Errorf("Server shutdown failed: %v", err)
	}

	wsServer.Close()

	pgPool.Close()
	logrus.Info("Сервер успешно завершил раоботу")
	return nil
}
