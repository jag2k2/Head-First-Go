package main

import (
	"robot/interfaces"
	"robot/transformers"
)

func main() {
	var robot interfaces.Robot = transformers.NewAutobot("Optimus Prime")
	robot.MakeSound()
	robot.Walk()
	robot = transformers.NewDecepticon("Starscream")
	robot.MakeSound()
	robot.Walk()
}
