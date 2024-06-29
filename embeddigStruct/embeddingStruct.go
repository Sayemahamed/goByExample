package main

import "fmt"

type base struct {
	num int
}
type container struct {
	base
	str string
}

func (b *base) describe() {
	fmt.Println("base with", b.num)
}
func main() {
	c := container{
		base{1},
		"hello",
	}
	fmt.Println(c)
	fmt.Println(c.str, c.num)
	fmt.Println(c.base.num)
	c.describe()
	c.base.describe()
}
