package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters [3]string = [3]string{"a", "b", "c"}
	floats := [2]float64{3.14, 1.414}
	
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
	fmt.Println(floats[0])
	fmt.Println(floats[1])
}