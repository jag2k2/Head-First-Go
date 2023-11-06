package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var response *http.Response // http.Response has a body field that represents the content of the page
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()                // http.Response satisfies ReadCloser interface which means it has a Close method and a Read method (see below)
	body, err := ioutil.ReadAll(response.Body) // returns ([]byte, error)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
