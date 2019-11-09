package repositories

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-microservice/src/api/domain/repositories"
	"go-microservice/src/api/services"
	"go-microservice/src/api/utils/errors"
	"go-microservice/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	funcCreateRepo func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(request []repositories.CreateRepoRequest)(repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct {
}

func (r *repoServiceMock) CreateRepo(clientId string,request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError){
	return funcCreateRepo(request)
}

func (r *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest)(repositories.CreateReposResponse, errors.ApiError){
	return funcCreateRepos(request)
}


func TestCreateRepNoErrorMockingTheEntireService(t *testing.T){
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (response *repositories.CreateRepoResponse, apiError errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id: 321,
			Name: "mocked service",
			Owner: "golang",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.Id)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang",result.Owner)
}

func TestCreateRepErrorFromGithubMockingTheEntireService(t *testing.T){
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (response *repositories.CreateRepoResponse, apiError errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}

