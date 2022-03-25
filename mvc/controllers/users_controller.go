package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/micro14/golang-microservices/mvc/services"
	"github.com/micro14/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	// Validate all the parameters that we need and we have
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// Return Bad Request to client
		userErr := &utils.ApiError{
			Message:   "user_id must be a member",
			ErrorCode: http.StatusBadRequest,
			Code:      "bad_request",
		}
		jsonVal, _ := json.Marshal(userErr)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonVal)
		return
	}
	// Validation End

	// Send to service for business logic
	user, userErr := services.GetUser(userId)
	if userErr != nil {
		// Handle the error and return to the client
		jsonVal, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.ErrorCode)
		resp.Write(jsonVal)
		return
	}

	// Return user to cloent
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
