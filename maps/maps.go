package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	nameMap := map[string][]string{
		"sayem": {"a", "b"},
		"fahim": {"c", "d"},
	}
	fmt.Println(nameMap["sayem"][0])
	fmt.Println(nameMap)
	fmt.Println(len(nameMap))
	delete(nameMap, "sayem")
	fmt.Println(nameMap)
	clear(nameMap)
	fmt.Println(nameMap)
	a := m["a"]
	fmt.Println(a)
	x := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(x)
	y := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(x, y) {
		fmt.Println("Equal")
	}
}
