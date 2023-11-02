package transformers

import "fmt"

type Decepticon struct {
	name string
}

func NewDecepticon(name string) *Decepticon {
	return &Decepticon{name: name}
}

func (decepticon Decepticon) MakeSound() {
	fmt.Println(decepticon.name, "- Get the Energon cubes!")
}
func (decepticon Decepticon) Walk() {
	fmt.Println(decepticon.name, "- Klunk Kloonk")
}
