package services

import (
	"go-microservice/mvc/domain"
	"go-microservice/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64)(*domain.User, *utils.ApplicationError){
	user, err := domain.UserDao.GetUser(userId)
	if err != nil{
		return nil, err
	}
	return user, err
}
