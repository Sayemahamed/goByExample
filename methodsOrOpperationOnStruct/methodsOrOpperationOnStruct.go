package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}
func (r *rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}
func main() {
	a := rectangle{width: 10, height: 5}
	fmt.Println(a.area())
	fmt.Println(a.perimeter())
	b := &a
	b.height = 11
	b.width = 11
	fmt.Println(a.area())
	fmt.Println(a.perimeter())
}
