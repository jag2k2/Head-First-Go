package main

import (
	"robot/interfaces"
	"robot/transformers"
)

func main() {
	var robot interfaces.Robot = transformers.NewAutobot("Optimus Prime")
	doRobotStuff(robot)
	var robot2 interfaces.Robot = transformers.NewDecepticon("Starscream")
	doRobotStuff(robot2)
}

func doRobotStuff(robot interfaces.Robot) {
	robot.MakeSound()
	robot.Walk()
}
