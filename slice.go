package main

import "fmt"


func main() {
	words := make([] string, 13)
	numbers := []int{2,3,4,5}
	words[0] = "on"
	words[1] = "go"
	words[2] = "make"
	words[3] = "eat"
	words[4] = "run"
	words[5] = "gone"
	words[6] = "jump"
	words[12] = "eye"
	fmt.Println(words)
	fmt.Println(numbers)
}