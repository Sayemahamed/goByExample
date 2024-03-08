package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	width  float64
	height float64
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r *rectangle) area() float64 {
	return r.width * r.height
}
func (c *circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}
func (r *rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {
	a := circle{2}
	b := rectangle{3, 4}
	measure(&a)
	measure(&b)
}
