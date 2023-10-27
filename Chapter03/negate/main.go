package main

import (
	"fmt"
)

func main() {
	var truth, lies bool = true, false
	negate(&truth)
	fmt.Println(truth)
	negate(&lies)
	fmt.Println(lies)
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}