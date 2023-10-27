package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year() // methods are functions associated with a particular type
	fmt.Println(year)
}
