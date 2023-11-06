package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	url  string
	size int
}

func responseSize(urlChan chan string, pageChan chan Page) {
	url := <-urlChan
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	pageChan <- Page{url: url, size: len(body)}
}

func main() {
	pageChan := make(chan Page)
	urlChan := make(chan string)

	// go responseSize("https://example.com", sizes)
	// go responseSize("https://golang.org", sizes)
	// go responseSize("https://golang.org/docs", sizes)
	// fmt.Println(<-sizes)
	// fmt.Println(<-sizes)
	// fmt.Println(<-sizes)

	urls := []string{"https://example.com/", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(urlChan, pageChan)
		urlChan <- url
	}
	for i := 0; i < len(urls); i++ {
		page := <-pageChan
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}
