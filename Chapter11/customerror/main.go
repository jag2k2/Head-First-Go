package main

import (
	"fmt"
	"log"
)

type OverheatError float64

func (overheatError OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", overheatError)
}

func main() {
	var err error = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}
