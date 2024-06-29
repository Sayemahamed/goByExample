package main

import "fmt"

func change(number int) {
	number = 69
}
func changeByPointer(number *int) {
	*number = 69
}
func main() {
	number := 10
	fmt.Println(number)
	change(number)
	fmt.Println(number)
	changeByPointer(&number)
	fmt.Println(number)
	fmt.Println(&number)
}
