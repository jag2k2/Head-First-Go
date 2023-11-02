package main

import (
	"calendar/day"
	"fmt"
	"log"
)

func main() {
	event := day.Event{}
	err := event.SetTitle("Jeffs Birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}
