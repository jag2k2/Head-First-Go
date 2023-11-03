package main

import "fmt"

type Sizeable interface {
	Size() int
}

type Node struct {
	label string
	props map[string]string
}

func NewNode(label string) *Node {
	return &Node{label: label, props: map[string]string{}}
}

func NewGraph() *Graph {
	return &Graph{nodes: make([]*Node, 0, 10)}
}

func (n *Node) Add(key string, val string) {
	n.props[key] = val
}

func (g *Graph) Size() int {
	return len(g.nodes)
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) Add(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) Sum() int {
	sum := 0
	for _, n := range g.nodes {
		sum += len(n.props)
	}
	return sum
}

func (g *Graph) Psum(ch chan int) {
	sum := 0
	for _, n := range g.nodes {
		sum += len(n.props)
	}
	ch <- sum
}

func main() {
	n1 := NewNode("today")
	n2 := NewNode("tomorrow")

	fmt.Println(n2)
	fmt.Printf("%T\n", n2)
	fmt.Printf("%v\n", n2)

	n1.Add("a", "1")

	n3 := Node{label: "not a pointer", props: map[string]string{}}
	n3.Add("b", "2")
	fmt.Println(n3)

	g := NewGraph()
	g.Add(n1)
	g.Add(n2)
	g.Add(&n3)
	fmt.Printf("%v\n", g)

	// var g2 *Graph = nil
	// var s Sizeable = g2
	// if s != nil { // s is NOT nil because it is assigned to g2.  However the interface value is nil!
	// 	fmt.Printf("%v\n", s.Size())
	// }

	ch := make(chan int)
	go g.Psum(ch)
	sum := <-ch
	fmt.Printf("%d\n", sum)
}
