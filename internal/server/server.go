package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Message struct {
	Text string `json:"msg"`
}

var messages []Message
var mu sync.Mutex

func Serv() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			sendHandler(w, r)
		case http.MethodGet:
			receiveHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe(":8010", mux))
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Failed to decode message", http.StatusBadRequest)
		return
	}

	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func receiveHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	if len(messages) == 0 {
		mu.Unlock()
		http.Error(w, "No messages available", http.StatusNotFound)
		return
	}

	msg := messages[0]
	messages = messages[1:]
	mu.Unlock()

	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
		return
	}
}
