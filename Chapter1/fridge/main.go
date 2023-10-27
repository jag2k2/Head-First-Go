package main

import(
	"fmt"
)

func main(){
	var originalCount int		//explicit variable declaraion
	eatenCount := 4				// short varialbe declaration.  var is an int because that is the type it is initialized with.
	originalCount = 10
	fmt.Println("I started with", originalCount, "apples.")			// functions names that begin with a capital letter are considered "exported"
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}