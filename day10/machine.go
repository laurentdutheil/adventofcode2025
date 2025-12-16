package day10

import "strings"

type Machine struct {
	Indicator string
	Buttons   []Button
}

func NewMachine(indicator string) *Machine {
	return &Machine{Indicator: indicator}
}

func (m *Machine) AddButton(b Button) {
	m.Buttons = append(m.Buttons, b)
}
func (m *Machine) FewestButtonPresses() int {
	indicator := strings.Repeat(".", len(m.Indicator))
	root := &Node{
		indicator: indicator,
		buttons:   []Button{},
		children:  []*Node{},
	}
	root.addChildren(m.Buttons)
	return m.breadthFirstSearch(root)

}

func (m *Machine) breadthFirstSearch(root *Node) int {
	var queue []*Node
	queue = append(queue, root.children...)

	for len(queue) > 0 {
		child := queue[0]
		queue = queue[1:]
		if child.indicator == m.Indicator {
			return len(child.buttons)
		}
		child.addChildren(m.Buttons)
		queue = append(queue, child.children...)
	}

	return 0
}
