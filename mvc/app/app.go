package app

import (
	"go-microservice/mvc/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}