package domain

import (
	"fmt"
	"net/http"

	"github.com/micro14/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Rishab", LastName: "Sinha", Email: "rishab12sinha@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApiError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApiError{
		Message:   fmt.Sprintf("User %v not found", userId),
		ErrorCode: http.StatusNotFound,
		Code:      "not_found",
	}
}
