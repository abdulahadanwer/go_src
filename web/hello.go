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

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func germanhHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Guten tag!")
}

func arabicHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salamalikum, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hallo", germanhHandler)
	http.HandleFunc("/salam", arabicHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
