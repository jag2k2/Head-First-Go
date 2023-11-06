package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("view.html")
	check(err)
	err = template.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8085", nil)
	log.Fatal(err)
}
