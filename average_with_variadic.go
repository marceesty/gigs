package main

import "fmt"

// Variadic returns an int, a float and several strings when called
func average(numbers ...float64) float64 {
	max := 0.00
	for _, number := range numbers{
		max += number
	}
	return max/float64(len(numbers))

}

func main() {
	fmt.Println(average(1, 23.12,30,0.001,50.2))

}
