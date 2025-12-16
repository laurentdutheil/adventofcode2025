package day10_test

import (
	. "adventofcode2025/day10"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachine_FewestButtonPresses(t *testing.T) {
	machine := NewMachine(".##.")
	machine.AddButton(NewButton([]int{3}))
	machine.AddButton(NewButton([]int{1, 3}))
	machine.AddButton(NewButton([]int{2}))
	machine.AddButton(NewButton([]int{2, 3}))
	machine.AddButton(NewButton([]int{0, 2}))
	machine.AddButton(NewButton([]int{0, 1}))

	assert.Equal(t, 2, machine.FewestButtonPresses())
}
