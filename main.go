package main

import (
	"fmt"
)

func main() {
	notes := [8]string{"do","re", "me", "fa", "so", "la", "ti", "do"}
	for i := 0; i <= 7; i++ {
		fmt.Println(i, notes[i])
	}
}