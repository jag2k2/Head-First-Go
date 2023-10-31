package main

import "fmt"
import "subscriber/magazine"

func main() {
	var subscriber1 *magazine.Subscriber = defaultSubscriber("Aman Singh")
	address := magazine.Address{Street: "123 Oak St.", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber1.HomeAddress = address
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	subscriber2.HomeAddress.Street = "321 Maple St."
	subscriber2.HomeAddress.City = "Frisco"
	subscriber2.HomeAddress.State = "TX"
	subscriber2.HomeAddress.PostalCode = "78713"
	printInfo(subscriber2)

	employee := magazine.Employee{Name: "Joy Carr", Salary: 60000}
	employee.Street = "456 Elm St."  	// Can omit inner struct name/type when using anonymous fields
	employee.City = "Bakersfield"
	employee.State = "California"
	employee.PostalCode = "85469"
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	fmt.Println(employee.Address)
}

func printInfo(sub *magazine.Subscriber) {
	fmt.Println("Name:", sub.Name)
	fmt.Println("Monthly rate:", sub.Rate)
	fmt.Println("Active?", sub.Active)
	fmt.Println("Address:", sub.HomeAddress)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var sub magazine.Subscriber = magazine.Subscriber{Name: name, Rate: 5.99, Active: true}
	return &sub
}

func applyDiscount(sub *magazine.Subscriber) {
	sub.Rate = 4.99
}