package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				done <- true
				fmt.Println("Received all the jobs")
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job:", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")
	<-done
	_, ok := <-jobs
	fmt.Println("Received new job:", ok)
}
