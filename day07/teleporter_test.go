package day07_test

import (
	. "adventofcode2025/day07"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstLine(t *testing.T) {
	teleporter := NewTeleporter()
	teleporter.FirstLine(".......S.......")

	assert.Equal(t, 1, teleporter.TachyonCount())
}

func TestNoSplit(t *testing.T) {
	teleporter := NewTeleporter()
	teleporter.FirstLine(".......S.......")
	teleporter.Line("...............")

	assert.Equal(t, 1, teleporter.TachyonCount())
}

func TestOneSplit(t *testing.T) {
	teleporter := NewTeleporter()
	teleporter.FirstLine(".......S.......")
	teleporter.Line("...............")
	teleporter.Line(".......^.......")

	assert.Equal(t, 2, teleporter.TachyonCount())
}

func TestTwoSplit(t *testing.T) {
	teleporter := NewTeleporter()
	teleporter.FirstLine(".......S.......")
	teleporter.Line("...............")
	teleporter.Line(".......^.......")
	teleporter.Line("...............")
	teleporter.Line("......^.^......")

	assert.Equal(t, 4, teleporter.TachyonCount())
}
