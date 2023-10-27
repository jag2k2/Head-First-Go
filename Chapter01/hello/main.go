package main // has to be main if you are running this code directly

import (
	"fmt" // importing a package called fmt
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello", "Go!") // Println function is part of the fmt pacakge
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("sup"))

	var width, length float64 = 1.5, 3.5 // you can omit the type if you declare a variable and assign it on the same line
	var area float64 = width * length
	fmt.Println("Area is:", area, "square somethings")
}
