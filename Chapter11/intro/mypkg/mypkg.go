package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (myType MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (myType MyType) MethodWithParameter(num float64) {
	fmt.Println("MethodWithParameter called:", num)
}

func (myType MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (myType MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
