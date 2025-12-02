package day1_test

import (
	"testing"

	. "adventofcode2025/day1"

	"github.com/stretchr/testify/assert"
)

func TestParseLineRight(t *testing.T) {
	dial := NewDial()
	p := NewParser(dial)
	p.Parse("R23")
	assert.Equal(t, 73, dial.Position)
}
