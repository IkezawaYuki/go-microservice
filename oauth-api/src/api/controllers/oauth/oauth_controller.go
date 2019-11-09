package oauth

import (
	"github.com/gin-gonic/gin"
	"go-microservice/oauth-api/src/api/domain/oauth"
	"go-microservice/oauth-api/src/api/services"
	"go-microservice/src/api/utils/errors"
	"net/http"
)

func CreateAccessToken(c *gin.Context){
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil{
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context){
	token, err := oauth.GetAccessTokenByToken(c.Param("token_id"))
	if err != nil{
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}