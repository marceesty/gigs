package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)


func main() {
	files, err := ioutil.ReadDir("first_webapp")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else { 
			fmt.Println("File:", file.Name())
		}
	}
}