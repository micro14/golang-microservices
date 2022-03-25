package services

import (
	"github.com/micro14/golang-microservices/mvc/domain"
	"github.com/micro14/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApiError) {
	return domain.GetUser(userId)
}
