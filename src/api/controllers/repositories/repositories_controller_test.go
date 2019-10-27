package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-microservice/src/api/clients/restclient"
	"go-microservice/src/api/utils/errors"
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
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())

}


func TestCreateRepoErrorFromGithub(t *testing.T){
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c.Request = request

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
	//todo
}