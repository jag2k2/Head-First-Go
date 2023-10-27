package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./go2.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
