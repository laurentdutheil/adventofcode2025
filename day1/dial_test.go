package day1_test

import (
	"testing"

	. "adventofcode2025/day1"

	"github.com/stretchr/testify/assert"
)

func TestDialStartAt50(t *testing.T) {
	dial := NewDial()
	assert.Equal(t, 50, dial.Position)
}

func TestDialTurnRight(t *testing.T) {
	dial := NewDial()
	dial.TurnRight(23)
	assert.Equal(t, 73, dial.Position)
}

func TestDialTurnLeft(t *testing.T) {
	dial := NewDial()
	dial.TurnLeft(12)
	assert.Equal(t, 38, dial.Position)
}
