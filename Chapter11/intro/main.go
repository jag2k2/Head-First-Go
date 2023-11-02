package main

import (
	"fmt"
	"intro/mypkg"
)

func main() {
	var value mypkg.MyInterface = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
