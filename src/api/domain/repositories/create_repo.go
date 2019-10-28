package repositories

import "go-microservice/src/api/utils/errors"

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateRepoResponse struct {
	Id int64 `json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}

type CreateReposResponse struct {
	Response CreateRepoResponse `json:"repo"`
	Error errors.ApiError `json:"error"`
}