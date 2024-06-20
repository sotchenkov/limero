package response

import (
	"encoding/json"
	"net/http"
)

type QueueList struct {
	QueueNames []string `json:"queue_names"`
}

type QueueInfo struct {
	Name    string `json:"name"`
	Presize int    `json:"presize"`
	Size    int    `json:"size"`
	Head    int    `json:"head"`
	Tail    int    `json:"tail"`
	Count   int    `json:"count"`
}

type QueueDeleteResponse struct {
	OK   bool   `json:"ok"`
	Info string `json:"info"`
	Name string `json:"name"`
}

type QueueCreateResponse struct {
	OK      bool   `json:"ok"`
	Info    string `json:"info"`
	Name    string `json:"name"`
	Presize int    `json:"presize"`
}

type RootResponse struct {
	Limero  string `json:"limero"`
	Version string `json:"version"`
	License string `json:"license"`
	Author  string `json:"author"`
	Docs    string `json:"docs"`
}

type Ping struct {
	Ping string `json:"ping"`
}

type OK struct {
	OK bool `json:"ok"`
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
