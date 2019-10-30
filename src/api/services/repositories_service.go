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
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest)(repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init(){
	RepositoryService = &reposService{}
}


func (s *reposService) CreateRepo(input repositories.CreateRepoRequest)(*repositories.CreateRepoResponse, errors.ApiError){
	if err := input.Validate(); err != nil{
		return nil, err
	}

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

func (s *reposService) CreateRepos(request []repositories.CreateRepoRequest)(repositories.CreateReposResponse, errors.ApiError)  {
	input := make(chan repositories.CreateReponsitoriesResult)
	output := make(chan repositories.CreateReposResponse)

	// todo 53:14
	go s.handleRepoResults(input, output)

	for _, current := range request{
		go s.createRepoConcurrent(current, input)
	}

	result := <- output
	return result, nil
}

func (s *reposService) handleRepoResults(input chan repositories.CreateReponsitoriesResult, output chan repositories.CreateReposResponse){
	var results repositories.CreateReposResponse
	for incomingEvent := range input{
		repoResult := repositories.CreateReponsitoriesResult{
			Response:incomingEvent.Response,
			Error:incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateReponsitoriesResult){
	if err := input.Validate(); err != nil{
		output <- repositories.CreateReponsitoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(input)
	if err != nil{
		output <- repositories.CreateReponsitoriesResult{Error:err}
		return
	}

	output <- repositories.CreateReponsitoriesResult{Response: result}
}