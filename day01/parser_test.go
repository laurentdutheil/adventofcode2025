package day01_test

import (
	"testing"

	. "adventofcode2025/day01"

	"github.com/stretchr/testify/assert"
)

func TestParseLineRight(t *testing.T) {
	dial := NewDial()
	p := NewParser(dial)
	p.Parse("R23")
	assert.Equal(t, 73, dial.Position)
}

func TestParseLineLeft(t *testing.T) {
	dial := NewDial()
	p := NewParser(dial)
	p.Parse("L12")
	assert.Equal(t, 38, dial.Position)
}

func TestParseFile(t *testing.T) {
	dial := NewDial()
	p := NewParser(dial)

	p.ParseFile("input.txt")

	TheAnswer := 6099
	assert.Equal(t, TheAnswer, dial.ZeroPointingCounts)
}
