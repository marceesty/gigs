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

	for guessess := 0; guessess < 10; guessess++ {
		fmt.Println("You have", 10-guessess, "guessess left")
	

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
		} else {
			fmt.Println("What a nice guess!")
			break
		}
	}
	fmt.Println("Sorry you have no guess left, Game Over!")
}
