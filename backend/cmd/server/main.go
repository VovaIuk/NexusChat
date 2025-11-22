package main

import (
	"backend/internal/wsserver"

	"github.com/sirupsen/logrus"
)

const (
	addr = "0.0.0.0:8003"
)

func main() {
	wsServer := wsserver.NewWsServer(addr)
	logrus.Info("WS server started")
	if err := wsServer.Start(); err != nil {
		logrus.Errorf("Error with ws server: %v", err)
	}
}
