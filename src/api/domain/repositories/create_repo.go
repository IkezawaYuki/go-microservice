package repositories

import (
	"go-microservice/src/api/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.ApiError{
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == ""{
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	Id int64 `json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status_code"`
	Results    []CreateReponsitoriesResult `json:"result"`
}

type CreateReponsitoriesResult struct {
	Response *CreateRepoResponse `json:"response"`
	Error errors.ApiError `json:"error"`
}