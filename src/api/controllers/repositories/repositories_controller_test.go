package repositories

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-microservice/src/api/clients/restclient"
	"go-microservice/src/api/domain/repositories"
	"go-microservice/src/api/utils/errors"
	"go-microservice/src/api/utils/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}


func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())

}


func TestCreateRepoErrorFromGithub(t *testing.T){
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Response: &http.Response{
			StatusCode:http.StatusUnauthorized,
			Body:ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())

}

func TestCreateRepNoError(t *testing.T){
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:"https://api.github.com/user/repos",
		HttpMethod:http.MethodPost,
		Response:&http.Response{
			StatusCode:http.StatusCreated,
			Body:ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "",result.Owner)
}









