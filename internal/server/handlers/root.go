package handlers

import (
	"net/http"

	"github.com/sotchenkov/limero/internal/lib/response"
)

// @Summary		Limero information
// @Description Returns information about the limero
// @Tags 		root
// @Accept 		json
// @Produce 	json
// @Success 	200  {object}  response.RootResponse
// @Failure     405  {object}  response.Error
// @Router      / [get]
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Send(w, http.StatusMethodNotAllowed, response.Error{Error: "method_not_allowed", Info: "Only GET methods are allowed"})
		return
	}

	response.Send(w, http.StatusOK, response.RootResponse{
		Limero:  "Welcome!",
		Version: "0.1",
		License: "MIT license",
		Author:  "Alexey Sotchenkov",
		Docs:    r.Host + "/swagger/",
	})
}

// @Summary		Limero information
// @Description Returns information about the limero
// @Tags 		root
// @Accept 		json
// @Produce 	json
// @Success 	200  {object}  response.Ping
// @Failure     405  {object}  response.Error
// @Router      /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Send(w, http.StatusMethodNotAllowed, response.Error{Error: "method_not_allowed", Info: "Only GET methods are allowed"})
		return
	}

	response.Send(w, http.StatusOK, response.Ping{Ping: "pong"})
}
