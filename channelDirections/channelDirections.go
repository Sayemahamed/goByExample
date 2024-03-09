package main

import "fmt"

func functionA(A chan<- string, msg string) {
	A <- msg
}
func functionB(B chan<- string, A <-chan string) {
	B <- <-A
}
func main() {
	a := make(chan string, 1)
	b := make(chan string, 1)
	functionA(a, "hello")
	functionB(b, a)
	fmt.Println(<-b)
}
