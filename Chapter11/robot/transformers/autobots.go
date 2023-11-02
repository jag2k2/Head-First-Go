package transformers

import "fmt"

type Autobot struct {
	name string
}

func NewAutobot(name string) *Autobot {
	return &Autobot{name: name}
}

func (autobot Autobot) MakeSound() {
	fmt.Println(autobot.name, "- Transform and Rollout!")
}
func (autobot Autobot) Walk() {
	fmt.Println(autobot.name, "- Klink Klank")
}
