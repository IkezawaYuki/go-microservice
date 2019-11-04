package oauth

import (
	"fmt"
	"go-microservice/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = fmt.Sprintf("USER_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string)(*AccessToken, errors.ApiError){
	token := tokens[accessToken]
	if token != nil{
		return nil, errors.NewNotFoundError("no access token found with given parameters")
	}
	return token, nil
}