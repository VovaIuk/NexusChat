package main

import (
	"backend/config"
	"backend/internal/adapter/postgres"
	"backend/internal/chat/get_chatheaders"
	"backend/internal/chat/get_chathistory"
	getchats "backend/internal/chat/get_chats"
	httpcontroller "backend/internal/controller/http"
	"backend/internal/user/login_user"
	"backend/internal/user/register_user"
	"backend/internal/wsserver"
	"backend/pkg/httpserver"
	jwttoken "backend/pkg/jwt_token"
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

	jwtManager := jwttoken.NewJWTManager(c.JWT)
	wsServer := wsserver.NewWSServers()
	wsServer.Start()

	login_user.New(pgPool, jwtManager)
	register_user.New(pgPool)

	getchats.New(pgPool)
	get_chathistory.New(pgPool, jwtManager)
	get_chatheaders.New(pgPool, jwtManager)

	router := httpcontroller.Router(wsServer, jwtManager)
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
