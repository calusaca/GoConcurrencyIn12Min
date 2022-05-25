package main

import (
	"fmt"	
)

func main() {
	c := make(chan string, 2) //buffer channel not needed someone to read before sending 
	c <- "hello"
	c <- "world"

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
	
}

