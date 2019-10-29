package repositories

import "go-microservice/src/api/utils/errors"

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.ApiError{

	// todo

	return nil
}

type CreateRepoResponse struct {
	Id int64 `json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status_code"`
	Result    []CreateResponsitoriesResult `json:"result"`
}

type CreateResponsitoriesResult struct {
	Response CreateRepoResponse `json:"response"`
	Error errors.ApiError `json:"error"`
}