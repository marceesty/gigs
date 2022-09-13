package main

import 

// Variadic returns an int, a float and several strings when called
func variadic(num int, numb float64, letters ... string) (int, float64, string, error) {
	return num, number, letters, error

}

func main() {
	num1, numb2, aplphabet, _ := variadic(2, 3.6, "a", "b", "c")
	fmt.Println(num1, numb2, aplphabet)
}