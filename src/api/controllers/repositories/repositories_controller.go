package repositories

import (
	"github.com/gin-gonic/gin"
	"go-microservice/src/api/domain/repositories"
	"go-microservice/src/api/services"
	"go-microservice/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context){
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientId := c.GetHeader("x-Client-Id")

	result, err := services.RepositoryService.CreateRepo(clientId, request)
	if err != nil{
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}


func CreateRepos(c *gin.Context){
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil{
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}