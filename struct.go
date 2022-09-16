package main

import "fmt"

var subcriber struct {
	name string
	rate float64
	active bool
}

func main() {
	subcriber.name = "Unachuks Nedu"
	subcriber.rate = 4.99
	subcriber.active = true

	fmt.Println("Name:", subcriber.name)
	fmt.Println("Monthly rate:", subcriber.rate)
	fmt.Println("Active?", subcriber.active)

}