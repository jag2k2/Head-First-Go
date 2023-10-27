package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin) //function
	input, err := reader.ReadString('\n')                //method
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade == 100 {
		status = "Perfect!"
	} else if grade >= 60 {
		status = "You pass!"
	} else {
		status = "You fail!"
	}
	fmt.Println(status)
}
