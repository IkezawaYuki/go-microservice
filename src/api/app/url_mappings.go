package app

import (
	"go-microservice/src/api/controllers/polo"
	"go-microservice/src/api/controllers/repositories"
)

func mapUrls(){
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}