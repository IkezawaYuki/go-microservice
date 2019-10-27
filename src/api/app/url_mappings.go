package app

import (
	"go-microservice/src/api/controllers/polo"
	"go-microservice/src/api/controllers/repositories"
)

func mapUrls(){
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}