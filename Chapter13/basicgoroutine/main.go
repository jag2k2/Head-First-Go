package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond)
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond)
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")
}
