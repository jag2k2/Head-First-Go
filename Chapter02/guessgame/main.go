package main

import (
	"bufio"
	"fmt"	//These are import path.  The package name happens to match
	"log"
	"math/rand"	//This is the import path.  the package name is rand
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var seconds int64 = time.Now().Unix()
	rand.Seed(seconds)						// what the heck is this?  Changing some variable at the package level?
	var target int = rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.", "Chan you guess it?")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var input string
	var err error

	var maxGuesses int = 10
	var success bool = false
	for guesses:=0; guesses < maxGuesses; guesses++ {
		fmt.Println("You have", maxGuesses - guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		var guess int
		guess, err = strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess == target{
			success = true
			fmt.Println("You did it!  Congrats!! The number was", target)
			break;
		} else if guess < target {
			fmt.Println("Your guess was LOW.")
		} else {
			fmt.Println("Your guess was HIGH.")
		}
	}
	if !success {
		fmt.Println("Sorry you are out of guesses.  The numbers was", target)
	}
}
