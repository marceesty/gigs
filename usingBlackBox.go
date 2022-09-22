//how to use interface in golang?
package main

import (
    "fmt"
)

type Walker interface {
    Walk() string
}

type Human string
type Dog string

func (human Human) Walk() string { //A human is a walker
    return "I'm a man and I walked!"
}

func (dog Dog) Walk() string { //A dog is a walker
    return "I'm a dog and I walked!"
}

//Make a walker walk
func MakeWalk(w Walker) {
    fmt.Println(w.Walk())
}

func main() {
    var human Human
    var dog Dog
    MakeWalk(human)
    MakeWalk(dog)
}

