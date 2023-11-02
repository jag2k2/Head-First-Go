package main

import (
	"calendar/day"
	"fmt"
	"log"
)

func main() {
	date := day.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	event := day.Event{}
	event.SetTitle("Jeffs Birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}
