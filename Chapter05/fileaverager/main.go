package main

import "fmt"
import "log"
import "fileaverager/readfile"
import "fileaverager/averager"

func main(){
	filename := "data.txt"
	numbers, err := readfile.ReadFloats(filename)
	if err != nil{
		log.Fatal(err)
	}
	average := averager.FloatAverage(numbers)
	fmt.Println(average)
}