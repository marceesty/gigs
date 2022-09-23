package main

import (
	. "fmt"
)

func count() {
	for i := 0; i < 10; i++ {
		Printf("%d\t", i)
	}
}

func main() {
	defer count()
	Println("how are you")
}