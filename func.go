package main
import (
	"fmt"
)
func paintCalculator(lenght int, width int) {
	area := length * width
	fmt.Printf("The required paint is %.3f", area)
	return 
}


func main() {
	paintCalculator(20, 4)
}