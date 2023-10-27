package main

import (
	"strings"
	"fmt"
)

func main() {
	var broken string = "G# r#cks!"
	var replacer *strings.Replacer = strings.NewReplacer("#", "o")	// this is a function because it belongs to a package (strings)
	var fixed string = replacer.Replace(broken)	// this is a method because it belongs to a value (replacer)
	fmt.Println(fixed)
}