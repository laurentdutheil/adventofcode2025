package day10_test

import (
	. "adventofcode2025/day10"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseOneLine(t *testing.T) {
	factory := NewFactory()
	parser := NewParser(factory)

	parser.ParseLine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")

	machine := factory.Machines[0]
	assert.Equal(t, ".##.", machine.Indicator)
	assert.Equal(t, []int{3}, machine.Buttons[0].LightsPosition)
	assert.Equal(t, []int{1, 3}, machine.Buttons[1].LightsPosition)
	assert.Equal(t, []int{2}, machine.Buttons[2].LightsPosition)
	assert.Equal(t, []int{2, 3}, machine.Buttons[3].LightsPosition)
	assert.Equal(t, []int{0, 2}, machine.Buttons[4].LightsPosition)
	assert.Equal(t, []int{0, 1}, machine.Buttons[5].LightsPosition)
}

func TestParser_ParseLine(t *testing.T) {
	factory := NewFactory()
	parser := NewParser(factory)

	parser.ParseLine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
	parser.ParseLine("[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}")
	parser.ParseLine("[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}")

	assert.Equal(t, 7, factory.FewestButtonPresses())
}

func TestParser_ParseFile(t *testing.T) {
	factory := NewFactory()
	parser := NewParser(factory)

	parser.ParseFile()

	assert.Equal(t, 432, factory.FewestButtonPresses())

}
