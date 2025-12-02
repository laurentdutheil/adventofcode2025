package day1_test

import (
	"testing"

	. "adventofcode2025/day1"

	"github.com/stretchr/testify/assert"
)

func TestDialStartAt50(t *testing.T) {
	dial := NewDial()
	assert.Equal(t, 50, dial.Position)
	assert.Equal(t, 0, dial.ZeroPointingCounts)
}

func TestDialTurnRight(t *testing.T) {
	dial := NewDial()
	dial.TurnRight(23)
	assert.Equal(t, 73, dial.Position)
	assert.Equal(t, 0, dial.ZeroPointingCounts)
}

func TestDialTurnLeft(t *testing.T) {
	dial := NewDial()
	dial.TurnLeft(12)
	assert.Equal(t, 38, dial.Position)
	assert.Equal(t, 0, dial.ZeroPointingCounts)
}

func TestDialTurnRightPassing0(t *testing.T) {
	dial := NewDial()
	dial.TurnRight(51)
	assert.Equal(t, 1, dial.Position)
	assert.Equal(t, 1, dial.ZeroPointingCounts)
}

func TestDialTurnLeftPassing0(t *testing.T) {
	dial := NewDial()
	dial.TurnLeft(51)
	assert.Equal(t, 99, dial.Position)
	assert.Equal(t, 1, dial.ZeroPointingCounts)
}

func TestDialPointToZeroTurningRight(t *testing.T) {
	dial := NewDial()
	dial.TurnRight(50)
	assert.Equal(t, 0, dial.Position)
	assert.Equal(t, 1, dial.ZeroPointingCounts)
}

func TestDialPointToZeroTurningLeft(t *testing.T) {
	dial := NewDial()
	dial.TurnLeft(50)
	assert.Equal(t, 0, dial.Position)
	assert.Equal(t, 1, dial.ZeroPointingCounts)
}

func TestDialPointToZeroTurningRight1000(t *testing.T) {
	dial := NewDial()
	dial.TurnRight(1000)
	assert.Equal(t, 50, dial.Position)
	assert.Equal(t, 10, dial.ZeroPointingCounts)
}

func TestDialPointToZeroTurningLeft1000(t *testing.T) {
	dial := NewDial()
	dial.TurnLeft(1000)
	assert.Equal(t, 50, dial.Position)
	assert.Equal(t, 10, dial.ZeroPointingCounts)
}

func TestSeveralTurns(t *testing.T) {
	dial := NewDial()
	dial.TurnLeft(68)
	dial.TurnLeft(30)
	dial.TurnRight(48)
	dial.TurnLeft(5)
	dial.TurnRight(60)
	dial.TurnLeft(55)
	dial.TurnLeft(1)
	dial.TurnLeft(99)
	dial.TurnRight(14)
	dial.TurnLeft(82)
	assert.Equal(t, 32, dial.Position)
	assert.Equal(t, 6, dial.ZeroPointingCounts)
}
