package main

import "fmt"

type Number int

func (number *Number) Add (otherNumber Number) {
	fmt.Println(*number, "plus", otherNumber, "is", *number + otherNumber)
}

func (number *Number) Subtract (otherNumber Number) {
	fmt.Println(*number, "minus", otherNumber, "is", *number - otherNumber)
}

func main() {
	ten := Number(10)
	four := Number(4)
	ten.Add(four)
	ten.Subtract(four)
}