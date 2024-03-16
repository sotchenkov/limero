package server

import (
	"log"
	"net/http"

	_ "github.com/sotchenkov/limero/docs"
	"github.com/sotchenkov/limero/internal/server/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Serv() {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:7920/swagger/doc.json"),
	))

	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/ping", handlers.Ping)
	mux.HandleFunc("/queue/", handlers.Queue)
	mux.HandleFunc("/msg", handlers.Msg)

	log.Fatal(http.ListenAndServe(":7920", mux))
}
