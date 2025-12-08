package day04_test

import (
	"adventofcode2025/day04"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessibleRolls(t *testing.T) {
	grid := day04.ParseFile()
	grid.MarkAccessibleRolls()
	assert.Equal(t, 1508, grid.CountMarkedRolls())
}

func TestRemoveAccessibleRolls(t *testing.T) {
	grid := day04.ParseFile()
	count := 0
	grid.MarkAccessibleRolls()
	countMarkedRolls := grid.CountMarkedRolls()
	for countMarkedRolls > 0 {
		count += countMarkedRolls
		grid.RemoveAccessibleRolls()
		grid.MarkAccessibleRolls()
		countMarkedRolls = grid.CountMarkedRolls()
	}

	assert.Equal(t, 8538, count)
}
