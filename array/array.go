package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println("initial array", array)
	array[4] = 100
	fmt.Println("set", array)
	fmt.Println("get", array[4])
	anotherArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array length :", len(anotherArray))
	fmt.Println("array", anotherArray)
	twoDArray := [][]int{{1, 2}, {3, 4}}
	fmt.Println("2d array", twoDArray)
	for i := 0; i < len(twoDArray); i++ {
		for j := 0; j < len(twoDArray[i]); j++ {
			fmt.Print(twoDArray[i][j], " ")
		}
		fmt.Println()
	}
}
