package main

import "fmt"

func val() (int, int) {
	return 6, 9
}
func main() {
	fmt.Println(val())
	a, b := val()
	fmt.Println(a)
	fmt.Println(b)
	c, _ := val()
	fmt.Println(c)
}
