package main

import "fmt"
import "subscriber/magazine"

func main() {
	var subscriber1 *magazine.Subscriber = defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}

func printInfo(sub *magazine.Subscriber) {
	fmt.Println("Name:", sub.Name)
	fmt.Println("Monthly rate:", sub.Rate)
	fmt.Println("Active?", sub.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var sub magazine.Subscriber = magazine.Subscriber{Name: name, Rate: 5.99, Active: true}
	return &sub
}

func applyDiscount(sub *magazine.Subscriber) {
	sub.Rate = 4.99
}