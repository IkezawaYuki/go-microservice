package controllers

import (
	"encoding/json"
	"go-microservice/mvc/services"
	"go-microservice/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request){
	userIdParam := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil{
		apiErr := &utils.ApplicationError{
			Message:"user_id must be a number",
			StatusCode:http.StatusBadRequest,
			Code:"bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil{
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}