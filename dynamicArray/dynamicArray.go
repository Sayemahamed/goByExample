package main

import "fmt"

func main() {
	var slice []string
	fmt.Println(slice, slice == nil, len(slice) == 0, cap(slice))
	slice = make([]string, 3)
	fmt.Println(slice, len(slice), cap(slice))
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println(slice)
	fmt.Println(slice[2])
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println(slice)
	aSlice := make([]string, len(slice))
	fmt.Println(aSlice)
	copy(aSlice, slice)
	fmt.Println(aSlice)
	l := slice[2:5]
	fmt.Println(l)
	l = slice[:3]
	fmt.Println(l)
	l = slice[4:]
	fmt.Println(l)
	twoDSlice := make([][2]int,2)
	fmt.Println(twoDSlice)
	for i := 0; i < len(twoDSlice); i++ {
		for j := 0; j < len(twoDSlice[i]); j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println(twoDSlice)
}
