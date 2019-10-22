package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}


func TestGetUserNoErr(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	//assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "YUKI", user.FirstName)
	assert.EqualValues(t, "IKEZAWA", user.LastName)
	assert.EqualValues(t, "nnn@aaa.com", user.Email)
	assert.EqualValues(t, 28, user.Age)


}