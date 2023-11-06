package main

import (
	"fmt"
	"time"
)

func greeting(myChannel chan string) {
	fmt.Println("Greeting initiated.  Will respond in 3 seconds:")
	time.Sleep(3 * time.Second)
	myChannel <- "hi"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println("Waiting for greeting:")
	response := <-myChannel
	fmt.Println(response)
}
