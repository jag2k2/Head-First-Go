package main

import (
	"log"
	"net/http"
)

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!\n")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut, web!\n")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!\n")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8085", nil)
	log.Fatal(err)
}
