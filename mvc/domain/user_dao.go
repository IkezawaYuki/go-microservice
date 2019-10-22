package domain

import (
	"fmt"
	"go-microservice/mvc/utils"
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
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
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
