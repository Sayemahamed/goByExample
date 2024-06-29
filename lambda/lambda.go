package main

import "fmt"

func main() {
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case float32:
			fmt.Println("I'm a float32")
		case int:
			fmt.Println("I'm a integer")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("I'm  mystery")
		}
	}
	whatAmI(1)
	whatAmI("sayem")
	whatAmI(1.5)
}
