package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	fmt.Println(str)
	fmt.Println("Rune count :", utf8.RuneCountInString(str), "length is :", len(str))
	for _, char := range str {
		fmt.Printf("%#U\n", char)
	}
}
