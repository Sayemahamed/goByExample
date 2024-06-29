package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	messenger := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			messenger <- i
		}
		close(messenger)
	}()
	for msg := range messenger {
		fmt.Println(msg)
	}
}
