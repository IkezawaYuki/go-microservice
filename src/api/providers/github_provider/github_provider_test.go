package github_provider

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-microservice/src/api/clients/restclient"
	"go-microservice/src/api/domain/github"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T){
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestAuthorizationHeader(t *testing.T){
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Err:errors.New("invalid resclient response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid resclient response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockups()

	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Response:&http.Response{
			StatusCode: http.StatusCreated,
			Body: invalidCloser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid resclient response", err.Message)

}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Response:&http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message: 1"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)

}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Response:&http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)

}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Response:&http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal github create repo response", err.Message)

}


func TestCreateNoError(t *testing.T){
	// todo
}