package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "Can You hear me on the other side?"
	}()
	msg := <-messages
	fmt.Println(msg)
}
