package main

import (
	"github.com/sotchenkov/limero/internal/server"
	"go.uber.org/zap"
)

// @title           Limero
// @version         0.1
// @description     This is a message broker

// @host      localhost:7920
// @license.name  MIT license
// @BasePath  /
func main() {
	zlog, _ := zap.NewProduction()

	server.Serv(zlog)
}
