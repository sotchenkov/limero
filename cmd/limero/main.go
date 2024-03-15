package main

import (
	"github.com/sotchenkov/limero/internal/server"
)

// @title           Limero
// @version         1.0
// @description     This is a message broker

// @host      localhost:7920
// @license.name  MIT license
// @BasePath  /

func main() {
	server.Serv()

}
