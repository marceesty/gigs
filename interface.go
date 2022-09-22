package main

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface{
	MakeSound()
}
//_____________________________________//

func (t TapeRecorder) play() {
	fmt.Println("playing...")
}

func (t TapeRecorder) stop() {
	fmt.Println("music stopped...")
}

func (t TapeRecorder) record() {
	fmt.Println("recording...")
}

func (t TapePlayer) play() {
	fmt.Println("playing music...")
}

func (t TapePlayer) stop() {
	fmt.Println("music stopped...")
}

type player interface {
	play()
	stop()
}

func playList(device player, songs []string) {
	for _, song := range songs {
		device.play(song)
	}
	device.stop()
}


func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	//__________________________________//

	var walkman player = TapePlayer{}
}