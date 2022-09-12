package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b"}
	slice = append(slice, "c", "d")
	fmt.Println(slice)
}