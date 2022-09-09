// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I'v chosen a random number,\n Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Make a guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops. Your guess is Low")
	} else if guess > target {
		fmt.Println("Oops. Your guess is High")
	}
	fmt.Println("My guess is: ", target)\
}