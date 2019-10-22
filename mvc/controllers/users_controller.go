package controllers

import (

	"github.com/gin-gonic/gin"
	"go-microservice/mvc/services"
	"go-microservice/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil{
		apiErr := &utils.ApplicationError{
			Message:"user_id must be a number",
			StatusCode:http.StatusBadRequest,
			Code:"bad_request",
		}
		utils.RespondErr(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)

	if apiErr != nil{
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}

