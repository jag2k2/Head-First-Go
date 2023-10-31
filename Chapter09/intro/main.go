package main

import "fmt"

type Person struct {
	name	string
	age		int
}

func (person Person) sayHi() {
	fmt.Println("Hi! My name is", person.name)
}

type Dog string

func (dog Dog) barkHi() {
	fmt.Println("Woof! My name is", dog)
}

func main() {
	person := Person{name: "Jeff", age: 43}
	person.sayHi()
	dog := Dog("Rusty")	//converts string to Dog because they both have same underlying string
	dog.barkHi()
}