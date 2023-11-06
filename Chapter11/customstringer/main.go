package main

import "fmt"

type Gallons float64

func (gallons Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", gallons)
}

type Liters float64

func (liters Liters) String() string {
	return fmt.Sprintf("%0.2f L", liters)
}

type Milliliters float64

func (milliliters Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", milliliters)
}

func main() {
	fmt.Println(Gallons(12.03248342))
	fmt.Println(Liters(12.03248342))
	fmt.Println(Milliliters(12.03248342))
}
