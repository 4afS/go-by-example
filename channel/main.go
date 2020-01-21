package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)

	go func() { messages <- "ping" }()
	go func() { messages <- "pong" }()
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
