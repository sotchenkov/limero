package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sotchenkov/limero/internal/lib/response"
	"github.com/sotchenkov/limero/internal/queue"
)

func Msg(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		PushMsg(w, r)
	case http.MethodGet:
		PopMsg(w, r)
	default:
		response.Send(w, http.StatusMethodNotAllowed, response.Error{Error: "method_not_allowed", Info: "Only PUT, DELETE methods"})
	}
}

// @Summary     Sends a message to queue
// @Description Sends a message to queue by name
// @Tags 		msg
// @Accept 		json
// @Produce 	json
// @Param       qname query     string  true  "Queue name"
// @Success 	201  {object}  response.OK
// @Failure     404  {object}  response.Error
// @Failure     400  {object}  response.Error
// @Router      /msg [post]
func PushMsg(w http.ResponseWriter, r *http.Request) {
	queueName := r.FormValue("qname")

	if queueName == "" {
		response.Send(w, http.StatusBadRequest, response.Error{Error: "missing_parameter", Info: "The queue name is required"})
		return
	}

	q, queueExist := queues[queueName]
	if !queueExist {
		response.Send(w, http.StatusNotFound, response.Error{Error: "not_found", Info: "A queue with this name was not found"})
		return
	}

	var msg *queue.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		response.Send(w, http.StatusBadRequest, response.Error{Error: "unsupported_type", Info: "The request body must be in json format"})
		return
	}

	if len(msg.Value) == 0 {
		response.Send(w, http.StatusBadRequest, response.Error{Error: "invalid_key", Info: "The request body must be JSON with the key 'value'"})
		return
	}

	q.Mu.Lock()
	q.Push(msg)
	q.Mu.Unlock()

	response.Send(w, http.StatusCreated, response.OK{OK: true})
}

// @Summary     Get message from the queue
// @Description Get message from the queue by name
// @Tags 		msg
// @Accept 		json
// @Produce 	json
// @Param       qname query     string  true  "Queue name"
// @Success 	200  {object}  queue.Message
// @Failure     404  {object}  response.Error
// @Failure     400  {object}  response.Error
// @Router      /msg [get]
func PopMsg(w http.ResponseWriter, r *http.Request) {
	queueName := r.FormValue("qname")

	if queueName == "" {
		response.Send(w, http.StatusBadRequest, response.Error{Error: "missing_parameter", Info: "The queue name is required"})
		return
	}

	q, queueExist := queues[queueName]
	if !queueExist {
		response.Send(w, http.StatusNotFound, response.Error{Error: "not_found", Info: "A queue with this name was not found"})
		return
	}

	q.Mu.Lock()

	if q.IsEmpty() {
		response.Send(w, http.StatusNotFound, response.Error{Error: "empty_queue", Info: "The queue is empty"})
		return
	}
	msg := q.Pop()

	q.Mu.Unlock()

	response.Send(w, http.StatusOK, msg)
}
