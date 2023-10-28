package main

import "fmt"

func main(){
	var notes []string = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes)
}