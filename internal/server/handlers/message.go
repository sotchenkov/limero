package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/sotchenkov/limero/internal/queue"
// )

// func SendHandler(w http.ResponseWriter, r *http.Request) {
// 	r.FormValue("name")
// 	var msg *queue.Message
// 	err := json.NewDecoder(r.Body).Decode(&msg)
// 	if err != nil {
// 		http.Error(w, "Failed to decode message", http.StatusBadRequest)
// 		return
// 	}

// 	mu.Lock()
// 	messages = append(messages, msg)
// 	q.Push(msg)
// 	mu.Unlock()

// 	w.WriteHeader(http.StatusOK)
// }

// func ReceiveHandler(w http.ResponseWriter) {
// 	mu.Lock()
// 	if len(messages) == 0 {
// 		mu.Unlock()
// 		http.Error(w, "No messages available", http.StatusNotFound)
// 		return
// 	}

// 	msg := messages[0]
// 	messages = messages[1:]
// 	mu.Unlock()

// 	msg := q.Pop()

// 	err := json.NewEncoder(w).Encode(msg)
// 	if err != nil {
// 		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
// 		return
// 	}
// }
