package main

import "fmt"

var firstBook string
func main() {
	firstBook = "Things fall apart"
	secondBook := &firstBook
	fmt.Println(*secondBook)
}