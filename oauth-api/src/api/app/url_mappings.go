package app

import (
	"go-microservice/oauth-api/src/api/controllers/oauth"
	"go-microservice/src/api/controllers/polo"
)

func mapUrls(){
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}