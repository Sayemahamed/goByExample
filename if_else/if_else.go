package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	var n int
	for {
		fmt.Println("Type 0 to quit")
		fmt.Println("Type a number: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		} else if n%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
