package oauth

import (
	"go-microservice/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? and password=?;"
)

var (
	users = map[string]*User{
		"Yuki": {Id: 123, Username:"Yuki"},
	}
)


func GetUserByUsernameAndPassword(username string, password string)(*User, errors.ApiError){
	user := users[username]
	if user == nil{
		return nil, errors.NewNotFoundError("no user found with given parameter")
	}
	return user, nil
}