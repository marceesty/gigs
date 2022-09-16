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
	HomeAddress Address

}

func main() {
	subscriber := Subscriber.Name["Kenyol"]
	fmt.Println(subscriber)
}

