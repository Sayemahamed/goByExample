package main

import (
	"fmt"
	"time"
)

func slowRoutine(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "slow"
		time.Sleep(time.Microsecond * 1800)
	}
	close(ch) // Close the channel when done sending
}

func fastRoutine(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "fast"
		time.Sleep(time.Microsecond * 100)
	}
	close(ch) // Close the channel when done sending
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go slowRoutine(ch1)
	go fastRoutine(ch2)
	for ch1 != nil || ch2 != nil {
		select {
		case msg1, ok1 := <-ch1:
			if !ok1 {
				ch1 = nil
			} else {
				fmt.Println(msg1)
			}
		case msg2, ok2 := <-ch2:
			if !ok2 {
				ch2 = nil
			} else {
				fmt.Println(msg2)
			}
		}
	}
}
