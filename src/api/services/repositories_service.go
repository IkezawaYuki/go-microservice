package services

import (
	"fmt"
	"go-microservice/src/api/config"
	"go-microservice/src/api/domain/github"
	"go-microservice/src/api/domain/repositories"
	"go-microservice/src/api/log"
	"go-microservice/src/api/providers/github_provider"
	"go-microservice/src/api/utils/errors"
	"net/http"
	"sync"

	"strings"
)

type reposService struct {}

type reposServiceInterface interface {
	CreateRepo(clientId string,request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest)(repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init(){
	RepositoryService = &reposService{}
}


func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest)(*repositories.CreateRepoResponse, errors.ApiError){
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
	log.Info("about to send request to external api", fmt.Sprintf("clientId:%s", clientId), "status:pending")
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil{
		log.Error("response obtained from external api", err, fmt.Sprintf("clientId:%s", clientId), "status:error")
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	log.Info("response obtained from external api", fmt.Sprintf("clientId:%s", clientId), "status:success")
	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner:response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest)(repositories.CreateReposResponse, errors.ApiError)  {
	input := make(chan repositories.CreateReponsitoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup

	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests{
		wg.Add(1)
		fmt.Println(current)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <- output

	successCreation := 0
	for _, current := range result.Results{
		if current.Response != nil{
			successCreation++
		}
	}
	if successCreation == 0{
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreation == len(requests){
		result.StatusCode = http.StatusCreated
	}else{
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateReponsitoriesResult, output chan repositories.CreateReposResponse){
	var results repositories.CreateReposResponse
	for incomingEvent := range input{
		repoResult := repositories.CreateReponsitoriesResult{
			Response:incomingEvent.Response,
			Error:incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
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

