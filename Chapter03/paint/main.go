package main

import (
	"fmt"
	"log"
)
//The simplest way to create an error is to pass a string to the errors package New function, which returns an error value 
var squareMetersPerLiter float64 = 10.0

func main() {
	var amount, total float64
	var err error

	amount, err = paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	
	fmt.Printf("%0.2f total needed\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("A width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("A height of %0.2f is invalid", height)
	}
	var area float64 =  width * height
	return area / squareMetersPerLiter, nil
}