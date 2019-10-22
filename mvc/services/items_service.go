package services

import (
	"go-microservice/mvc/domain"
	"go-microservice/mvc/utils"
	"net/http"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (s *itemService) GetItem(itemId string)(*domain.Item, *utils.ApplicationError){
	return nil, &utils.ApplicationError{
		Message:"implement me",
		StatusCode:http.StatusInternalServerError,
	}
}
