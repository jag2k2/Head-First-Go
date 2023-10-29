package main

import "fmt"
import "log"
import "voter/datafile"

func main() {
	contents, err := datafile.ReadStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var ranks map[string]int = make(map[string]int)
	for _, line := range contents {
		ranks[line]++
	}
	fmt.Println(ranks)
}

