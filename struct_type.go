package main

import "fmt"

type car struct {
	name string
	topSpeed float64
}

type part struct {
	description string
	count int
}
func double(number int) {
	result := number * 2
	fmt.Println(result)
} 

func main() {
	double(10)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 320
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count =  24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

}