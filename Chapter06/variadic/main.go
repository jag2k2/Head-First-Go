package main

import (
	"fmt"
)

func main() {
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}