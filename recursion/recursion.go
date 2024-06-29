package main

import "fmt"

func fibonacci(n int) int {
	if n < 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func factorial(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println(fibonacci(7))
	fmt.Println(factorial(7))
}
