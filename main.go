package main

import (
	"fmt"
)

func main() {
	var num int
	anotherNumber := &num

	num = 10
	fmt.Println(num)
	fmt.Println(*anotherNumber)
}
