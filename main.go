package main

import (
	"fmt"
)

func paintCalculator(length float64, width float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("put a positive number")
	} else if length < 0 {
		return 0, fmt.Errorf("put a positive number")
	}
	area := length * width
	fmt.Printf("The required paint is %.1f liters\n", area)
	return area, nil
}

func main() {
	amount, err := paintCalculator(1, -4)
	fmt.Println(err)
	fmt.Println(amount, err)
}
