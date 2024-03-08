package main

import (
	"errors"
	"fmt"
)

func canIWorkWith(age int) (int, error) {
	if age == 3 {
		return 0, errors.New("canIWorkWith=>3")
	}
	if age > 3 {
		if age, err := canIWorkWith(age - 1); err != nil {
			return age, errors.New("canIWorkWith->" + err.Error())
		}
	}
	return age + 2, nil
}
func main() {
	if age, err := canIWorkWith(6); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(age)
	}
	if age, err := canIWorkWith(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(age)
	}
}
