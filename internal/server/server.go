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
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodPost:
	// 		handlers.SendHandler(w, r)
	// 	case http.MethodGet:
	// 		handlers.ReceiveHandler(w)
	// 	default:
	// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 	}
	// })

	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:7920/swagger/doc.json"),
	))

	mux.HandleFunc("/queue/", handlers.Queue)
	log.Fatal(http.ListenAndServe(":7920", mux))
}
