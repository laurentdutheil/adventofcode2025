package day10

import (
	"fmt"
)

type Factory struct {
	Machines []*Machine
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) AddMachin(m *Machine) {
	f.Machines = append(f.Machines, m)
}

func (f *Factory) FewestButtonPresses() int {
	result := 0
	for i, machine := range f.Machines {
		fmt.Printf("%d/%d\n", i, len(f.Machines))
		result += machine.FewestButtonPresses()
	}
	return result
}

type Node struct {
	indicator string
	buttons   []Button
	children  []*Node
}

func (node *Node) addChildren(buttons []Button) {
	for _, button := range buttons {
		r := button.Toggle(node.indicator)
		n := &Node{
			indicator: r,
			buttons:   append(node.buttons, button),
			children:  []*Node{},
		}
		node.children = append(node.children, n)
	}
}
