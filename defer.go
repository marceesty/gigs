package main

import (
	"fmt"
)

func main() {
	for i := 0; 1 < 10; i++ {
		defer fmt.Printf("%d", i)
	}
}