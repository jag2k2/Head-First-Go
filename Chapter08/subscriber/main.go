package main

import "fmt"

type subscriber struct {
	name 	string
	rate 	float64
	active 	bool
}

func printInfo(sub subscriber) {
	fmt.Println("Name:", sub.name)
	fmt.Println("Monthly rate:", sub.rate)
	fmt.Println("Active?", sub.active)
}

func defaultSubscriber(name string) subscriber {
	var sub subscriber
	sub.name = name
	sub.rate = 5.99
	sub.active = true
	return sub
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}