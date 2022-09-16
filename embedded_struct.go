package main

import "fmt"

type Address struct {
	Street string
	Number string
	Zipcode string
}

type Subscriber struct {
	Name string
	Phone int
	Active bool
	Address

}

func main() {
	subscriber := Subscriber{Name: "Kenyol"}
	subscriber.Number = "20"
	fmt.Println(subscriber)

	//Deleting homeAddress above and having only address still make address accessible
	
}