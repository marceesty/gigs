package main

import (
	"fmt"
)

func greeting(myChannel chan string) {
	myChannel <- "Hi How Are You?"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	// assigning chan to a varible
	received := <-myChannel
	fmt.Println(<-myChannel)

	// assigning chan to a varible
	
	fmt.Println(received)
}