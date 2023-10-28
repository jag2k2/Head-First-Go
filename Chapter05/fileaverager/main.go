package main

import "fmt"
import "log"
import "fileaverager/readfile"

func main(){
	filename := "data.txt"
	numbers, err := readfile.ReadFloats(filename)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(numbers)
}