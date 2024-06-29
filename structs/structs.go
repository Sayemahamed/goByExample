package main

import "fmt"

type person struct {
	name        string
	age         int
	personality string
}

func newPerson(name string, age int, personality string) *person {
	return &person{name, age, personality}
}
func main() {
	sayem := newPerson("sayem", 23, "Giga Chad")
	fmt.Println(sayem)
	fmt.Println(sayem.name, sayem.age, sayem.personality)
	thinkTank := person{"thinkTank", 23, "also known as sayem"}
	fmt.Println(thinkTank)
	sayem = &thinkTank
	fmt.Println(sayem)
}
