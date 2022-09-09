package main

import (
	"fmt"
)

func main() {
	numbers := [6]float64{0, 6, 10, -2, 40, 51}
	for i, number := range numbers {
		if number >= 5 && i < 55 {
		fmt.Println(i, number)
	}
	}
}