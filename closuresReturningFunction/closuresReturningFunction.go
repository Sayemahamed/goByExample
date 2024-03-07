package main

import "fmt"

func intFunc(flag int) func() int {
	i := 0
	if flag == 1 {
		return func() int {
			i++
			return i
		}
	} else if flag == 2 {
		return func() int {
			i--
			return i
		}
	} else if flag == 3 {
		return func() int {
			i *= 2
			return i
		}
	}
	return func() int {
		i /= 2
		return i
	}
}
func main() {
	add := intFunc(1)
	sub := intFunc(2)
	mul := intFunc(3)
	div := intFunc(0)
	fmt.Println(add()) //1
	fmt.Println(add()) //2
	fmt.Println(mul()) //0
	fmt.Println(mul()) //0
	fmt.Println(sub()) //-1
	fmt.Println(div()) //0
}
