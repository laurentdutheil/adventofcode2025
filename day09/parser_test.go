package day09_test

import (
	. "adventofcode2025/day09"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseLine(t *testing.T) {
	floor := &Floor{}
	parser := NewParser(floor)

	parser.ParseLine("7,1")
	parser.ParseLine("11,1")
	parser.ParseLine("11,7")
	parser.ParseLine("9,7")
	parser.ParseLine("9,5")
	parser.ParseLine("2,5")
	parser.ParseLine("2,3")
	parser.ParseLine("7,3")

	area := floor.LargestRectangle()
	assert.Equal(t, 50, area)
}

func TestParser_ParseFile(t *testing.T) {
	floor := &Floor{}
	parser := NewParser(floor)

	parser.ParseFile()
	area := floor.LargestRectangle()

	assert.Equal(t, 4777816465, area)
}
