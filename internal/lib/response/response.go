package response

import (
	"encoding/json"
	"net/http"
)

type QueueList struct {
	QueueNames []string `json:"queue_names"`
}

type QueueInfo struct {
	Name  string
	Size  int
	Head  int
	Tail  int
	Count int
}

type QueueDeleteResponse struct {
	OK   bool   `json:"ok"`
	Info string `json:"info"`
	Name string `json:"name"`
}

type QueueCreateResponse struct {
	OK   bool   `json:"ok"`
	Info string `json:"info"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

type Error struct {
	Error string `json:"error"`
	Info  string `json:"info"`
}

func JSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return jsonData, nil
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func Send(w http.ResponseWriter, statusCode int, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	setContentType(w)
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}
