package services

import (
	"go-microservice/mvc/domain"
	"go-microservice/mvc/utils"
)

func GetUser(userId int64)(*domain.User, *utils.ApplicationError){
	return domain.GetUser(userId)
}
