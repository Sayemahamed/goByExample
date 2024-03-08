package main

import (
	"fmt"
)

type userId int
type customData interface {
	~int | ~float32 | ~float64 | ~int64 | ~int32 | ~string
}
type user[T customData] struct {
	id   int
	name string
	data T
}

func add[T ~int | ~float32 | ~float64 | ~int64 | ~int32](a T, b T) T {
	return a + b
}
func mutator[T ~int | ~float32 | ~float64 | ~int64 | ~int32](values []T, logic func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValues = append(newValues, logic(value))
	}
	return newValues
}
func main() {
	a := userId(1)
	b := userId(2)
	fmt.Println(add(a, b))
	fmt.Println(add(1.7, 2.3))
	fmt.Println(mutator([]float32{1, 2, 3}, func(value float32) float32 { return value * 2.5 }))
	sayem := user[int]{1, "sayem", 10}
	fmt.Println(sayem)
	fahim := user[string]{1, "fahim", "Giga Chad"}
	fmt.Println(fahim)
}
