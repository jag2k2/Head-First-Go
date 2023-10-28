package main

import "fmt"
import "os"
import "strconv"
import "average2/averager"

func main(){
	var stringNums []string = os.Args[1:]
	var floatNums []float64
	for i, str := range stringNums {
		number, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("Could not parse element %d: %v\n", i, err)
		} else {
			floatNums = append(floatNums, number)
		}
	}
	average := averager.FloatAverage(floatNums)
	fmt.Printf("Average: %0.2f\n", average)
}