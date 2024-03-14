package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1500 * time.Millisecond)
	ticker.Stop()
	quit <- true
	fmt.Println("Ticker stopped")
}
