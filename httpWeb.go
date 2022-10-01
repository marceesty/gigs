package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request)  {
	write(writer, "Hello web!")
}

func igboHandler(writer http.ResponseWriter, request *http.Request)  {
	write(writer, "Dalu web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request)  {
	write(writer, "Salut web!")
}
func main() {
	http.HandleFunc("/english", englishHandler)
	http.HandleFunc("/igbo", igboHandler)
	http.HandleFunc("/french", frenchHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}