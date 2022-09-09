package main

import (
	"fmt"
)


func main() {
	var evenNumbers = [4]int{2,4,6,8}
	var myArray [10] string
	notes := [8]string{"do","re", "me", "fa", "so", "la", "ti", "do"}
	myArray[0] = "Chuks"
	myArray[5] = "89"
	myArray[8] = "welldone"
	fmt.Println(myArray)
	fmt.Println(evenNumbers)
	fmt.Println(notes)
}