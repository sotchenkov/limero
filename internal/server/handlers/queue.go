package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sotchenkov/limero/internal/lib/response"
	"github.com/sotchenkov/limero/internal/queue"
)

var queues = make(map[string]*queue.Queue)

func Queue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		createQueue(w, r)
	case http.MethodDelete:
		deleteQueue(w, r)
	case http.MethodGet:
		fmt.Println(len(r.URL.Path[len("/queue/"):]))
		if len(r.URL.Path[len("/queue/"):]) > 0 {
			queueInfo(w, r)
		} else {
			Queues(w, r)
		}
	default:
		response.Send(w, http.StatusMethodNotAllowed, response.Error{Error: "method_not_allowed", Info: "Only PUT, DELETE methods"})
	}
}

// @Summary     Creates a new queue
// @Description Creates a new queue with a name and size
// @Tags 		queue
// @Accept 		json
// @Produce 	json
// @Param       name query     string  true  "Queue name"
// @Param       size query     int 	   false "Queue size"
// @Success 	201  {object}  response.QueueCreateResponse
// @Failure     400  {object}  response.Error
// @Failure     409  {object}  response.Error
// @Router      /queue [put]
func createQueue(w http.ResponseWriter, r *http.Request) {

	queueName := r.FormValue("name")
	queueSizeStr := r.FormValue("size")

	if queueName == "" {
		response.Send(w, http.StatusBadRequest, response.Error{Error: "missing_parameter", Info: "The queue name is required"})
		return
	}

	// Default size to 1 if not specified
	var queueSize int
	if queueSizeStr == "" {
		queueSize = 1
	} else {
		var err error
		queueSize, err = strconv.Atoi(queueSizeStr)
		if err != nil {
			response.Send(w, http.StatusBadRequest, response.Error{Error: "unsupported_type", Info: "Unsupported queue size"})
			return
		}
	}

	if _, exists := queues[queueName]; exists {
		response.Send(w, http.StatusConflict, response.Error{Error: "already_exist", Info: "Queue already exists"})
		return
	}

	newQueue := queue.NewQueue(queueSize, queueName)
	queues[queueName] = newQueue

	response.Send(w, http.StatusCreated, response.QueueCreateResponse{OK: true, Info: "The queue has been created", Name: queueName, Size: queueSize})
}

// @Summary     Deletes a queue
// @Description Deletes a queue by name
// @Tags 		queue
// @Accept 		json
// @Produce 	json
// @Param       name query     string  true  "Queue name"
// @Success 	200  {object}  response.QueueDeleteResponse
// @Failure     404  {object}  response.Error
// @Router      /queue [delete]
func deleteQueue(w http.ResponseWriter, r *http.Request) {

	queueName := r.FormValue("name")
	if queueName == "" {
		response.Send(w, http.StatusBadRequest, response.Error{Error: "missing_parameter", Info: "The queue name is required"})
		return
	}

	_, queueExist := queues[queueName]
	if !queueExist {
		response.Send(w, http.StatusNotFound, response.Error{Error: "not_found", Info: "A queue with this name was not found"})
		return
	}

	delete(queues, queueName)

	response.Send(w, http.StatusOK, response.QueueDeleteResponse{OK: true, Info: "The queue has been deleted", Name: queueName})
}

// @Summary		Information about queue
// @Description Returns information about the queue by its name
// @Tags 		queue
// @Accept 		json
// @Produce 	json
// @Success 	200  {object}  response.QueueInfo
// @Failure     400  {object}  response.Error
// @Failure     404  {object}  response.Error
// @Router      /queue/{qname} [get]
func queueInfo(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")

	if len(segments) < 2 || segments[2] == "" {

		response.Send(w, http.StatusBadRequest, response.Error{Error: "missing_parameter", Info: "The queue name is required"})
		return
	}

	_, queueExist := queues[segments[2]]
	if !queueExist {
		response.Send(w, http.StatusNotFound, response.Error{Error: "not_found", Info: "A queue with this name was not found"})
		return
	}

	q := queues[segments[2]]
	response.Send(w, http.StatusOK, q.Info())

}

// @Summary     Queue list
// @Description Returns a list of queue names
// @Tags 		queue
// @Accept 		json
// @Produce 	json
// @Success 	200  {object}  response.QueueList
// @Router      /queue [get]
func Queues(w http.ResponseWriter, r *http.Request) {

	queueNames := make([]string, 0)
	for queueName := range queues {
		queueNames = append(queueNames, queueName)
	}
	response.Send(w, http.StatusOK, response.QueueList{QueueNames: queueNames})
}
