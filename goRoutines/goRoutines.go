package main

import (
	"fmt"
	"time"
)

func aFunction(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, str)
	}
}
func main() {
	aFunction("sayem")
	go aFunction("Giga Chad")
	go func(msg string) {
		fmt.Println(msg)
	}("Think Tank")
	fmt.Println("Yo Dude")
	time.Sleep(time.Second)
	fmt.Println("Done")
}
