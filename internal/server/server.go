package server

import (
	"log"
	"net/http"

	"net/http/pprof"

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

	// Регистрация pprof-обработчиков
    mux.HandleFunc("/debug/pprof/", pprof.Index)
    mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
    mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/ping", handlers.Ping)
	mux.HandleFunc("/queue", handlers.ActionOnQueueHandlers)
	mux.HandleFunc("/queue/", handlers.InfoAboutQueueHandlers)
	mux.HandleFunc("/msg", handlers.Msg)

	loggedMux := logger.LoggerMiddleware(zlog)(mux)

	log.Fatal(http.ListenAndServe(":7920", loggedMux))
}
