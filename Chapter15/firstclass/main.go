package main

import "fmt"

func sayHi(name string) {
	fmt.Println("Hi", name)
}

func sayBye(name string) {
	fmt.Println("Bye", name)
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func multiply(a int, b int) float64 {
	return float64(a * b)
}

func twice(theFunction func(string), name string) {
	theFunction(name)
	theFunction(name)
}

func doMath(passedFunction func(int, int) float64) {
	result := passedFunction(10, 2)
	fmt.Println(result)
}

func main() {
	twice(sayHi, "Jeff")
	twice(sayBye, "Aubrey")
	var mathFunction func(int, int) float64
	mathFunction = divide
	fmt.Println(mathFunction(5, 2))
	doMath(mathFunction)
	mathFunction = multiply
	doMath(multiply)
}
