package main

import "fmt"

func main(){
	c := make(chan string)
	go func (input chan string){
		fmt.Println("sending to the channel")
		input <- "hello"
		input <- "hello"
		input <- "hello"
		input <- "hello"
		input <- "hello"

	}(c)

	fmt.Println("receiving from the channel")
	greeting := <- c

	fmt.Println("greeting received")

	fmt.Println(greeting)
}

func helloWorld(){
	fmt.Println("hello world")
}