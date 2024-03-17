package server

import (
	"log"
	"net/http"

	_ "github.com/sotchenkov/limero/docs"
	"github.com/sotchenkov/limero/internal/lib/logger"
	"github.com/sotchenkov/limero/internal/server/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func Serv(zlog *zap.Logger) {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:7920/swagger/doc.json"),
	))

	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/ping", handlers.Ping)
	mux.HandleFunc("/queue/", handlers.Queue)
	mux.HandleFunc("/msg", handlers.Msg)

	loggedMux := logger.LoggerMiddleware(zlog)(mux)

	log.Fatal(http.ListenAndServe(":7920", loggedMux))
}
