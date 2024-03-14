package main

func main() {
	queue := make(chan int, 5)
	queue <- 0
	queue <- 1
	queue <- 3
	queue <- 2
	queue <- 4
	close(queue)
	for data := range queue {
		println(data)
	}
}
