package services

import (
	"github.com/stretchr/testify/assert"
	"go-microservice/mvc/domain"
	"go-microservice/mvc/utils"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64)(*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{

}

func (m *usersDaoMock) GetUser(userId int64)(*domain.User, *utils.ApplicationError){
	return getUserFunction(userId)
}

func init(){
	domain.UserDao = &usersDaoMock{}
}

func TestUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:"user 0 does not exists",
		}
	}

	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return &domain.User{
			Id:123,
		}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
