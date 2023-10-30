package main

import "fmt"

type part struct {
	description	string
	count 		int
}

type car struct {
	name 		string
	topSpeed	float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)

	var bolts part = minimumOrder("Hex Bolts")
	showInfo(bolts)
}

func showInfo(prt part){
	fmt.Println("Desc:", prt.description)
	fmt.Println("Count:", prt.count)
}

func minimumOrder(description string) part {
	var prt part
	prt.description = description
	prt.count = 100
	return prt
}