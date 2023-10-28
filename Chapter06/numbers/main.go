package main

import "fmt"

func main(){
	var notes []string = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes)

	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string {"a", "b", "c"}
	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}