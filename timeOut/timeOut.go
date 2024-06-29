package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string, 1)
	channel2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		channel1 <- "Hello from func 1"
	}()
	go func() {
		time.Sleep(time.Second)
		channel2 <- "Hello from func 2"
	}()
	select {
	case response := <-channel1:
		{
			fmt.Println(response)
		}
	case <-time.After(time.Second * 1):
		{
			fmt.Println("Timed out")
		}
	}
	select {
	case response := <-channel2:
		{
			fmt.Println(response)
		}
	case <-time.After(time.Second * 2):
		{
			fmt.Println("Timed out")
		}
	}
}
