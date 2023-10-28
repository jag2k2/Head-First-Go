package main

import "fmt"
import "log"
import "voter/datafile"

func main() {
	contents, err := datafile.ReadStrings("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contents)
}

