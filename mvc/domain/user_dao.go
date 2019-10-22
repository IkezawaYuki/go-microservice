package domain

import (
	"fmt"
	"go-microservice/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        1,
			FirstName: "YUKI",
			LastName:  "IKEZAWA",
			Email:     "nnn@aaa.com",
			Age:       28,
		},
		111: {
			Id:        1,
			FirstName: "YUKI2",
			LastName:  "IKEZAWA2",
			Email:     "nnn@aaa.com2",
			Age:       29,
		},
	}

	UserDao userDaoInterface
)

func init(){
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	user := users[userId]
	if user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:"not_found",
	}

}
