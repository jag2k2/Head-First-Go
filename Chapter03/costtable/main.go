package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("-----------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 98)
	fmt.Println();
	fmt.Printf("%12s | %7.3f\n", "%7.3f", 12.3456)
	fmt.Printf("%12s | %7.2f\n", "%7.2f", 12.3456)
	fmt.Printf("%12s | %7.1f\n", "%7.1f", 12.3456)
	fmt.Printf("%12s | %.1f\n", "%.1f", 12.3456)
	fmt.Printf("%12s | %.2f\n", "%.2f", 12.3456)
}