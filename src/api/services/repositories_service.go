package services

import (
	"go-microservice/src/api/config"
	"go-microservice/src/api/domain/github"
	"go-microservice/src/api/domain/repositories"
	"go-microservice/src/api/providers/github_provider"
	"go-microservice/src/api/utils/errors"

	"strings"
)

type reposService struct {}

type reposServiceInterface interface {
}

var (
	RepositoryService reposServiceInterface
)

func init(){
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest)(*repositories.CreateRepoResponse, errors.ApiError){
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == ""{
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:input.Name,
		Description:input.Description,
		Private:false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil{
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner:response.Owner.Login,
	}

	return &result, nil
}