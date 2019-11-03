package app

import (
	"github.com/gin-gonic/gin"
	"go-microservice/src/api/log"
)

var (
	router *gin.Engine
)

func init(){
	router = gin.Default()
}

func StartApp(){
	log.Info("about to map the urls", "step:1", "status:spending")
	mapUrls()
	log.Info("urls successfully mapped", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil{
		panic(err)
	}
}