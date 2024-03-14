package main

func fib(num int) int {
	if num <= 1 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}
func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- fib(j)
	}
}
func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	go worker(jobs, result)
	for i := 2; i <= 20; i++ {
		jobs <- i
	}
	close(jobs)
	for data := range result {
		println(data)
	}
	close(result)
}
