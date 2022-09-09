package main

import (
	"fmt"
)

func main() {
	notes := [8]string{"do","re", "me", "fa", "so", "la", "ti", "do"}
	for index, value := range notes {
		fmt.Println(index, value)
	}
}