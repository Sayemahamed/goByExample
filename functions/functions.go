package main

import "fmt"

func plus(x int, y int) int {
	return x + y
}
func plusPlus(x int, y int, z int) int {
	return x + y + z
}
func main() {
	fmt.Println("1 + 2 =", plus(1, 2))          // "plus(1, 2))
	fmt.Println("1 + 2 + 3 =", plusPlus(1, 2, 3)) // "plusPlus(1, 2, 3))
}
