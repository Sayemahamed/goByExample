package main

import "fmt"

func sum(num ...int) int {
	temp := 0
	for _, n := range num {
		temp += n
	}
	return temp
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(nums...))
}
