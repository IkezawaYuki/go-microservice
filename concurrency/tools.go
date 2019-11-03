package main

import (
	"bufio"
	"fmt"
	"go-microservice/src/api/domain/repositories"
	"go-microservice/src/api/utils/errors"
	"os"
)

type createRepoResult struct {
	Result *repositories.CreateRepoResponse
	Error errors.ApiError
}

func getRequests()[]repositories.CreateRepoRequest{
	//result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/Users/yikezawa/Desktop/request.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
	}
	return nil
}


func main(){
	getRequests()
}






//func main(){
//	c := make(chan string)
//	go func (input chan string){
//		fmt.Println("sending to the channel")
//		input <- "hello"
//		input <- "hello"
//		input <- "hello"
//		input <- "hello"
//		input <- "hello"
//
//	}(c)
//
//	fmt.Println("receiving from the channel")
//	greeting := <- c
//
//	fmt.Println("greeting received")
//
//	fmt.Println(greeting)
//}
//
//func helloWorld(){
//	fmt.Println("hello world")
//}