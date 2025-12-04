package day4_test

import (
	. "adventofcode2025/day4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAdjacents(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		expected string
	}{
		{"adjacents of 5", Position{Column: 1, Row: 1}, "12346789"},
		{"adjacents of 1", Position{Column: 0, Row: 0}, "....2.45"},
		{"adjacents of 9", Position{Column: 2, Row: 2}, "56.8...."},
		{"adjacents of 7", Position{Column: 0, Row: 2}, ".45.8..."},
		{"adjacents of 3", Position{Column: 2, Row: 0}, "...2.56."},
	}
	grid := NewGrid()
	grid.AddLine("123")
	grid.AddLine("456")
	grid.AddLine("789")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, grid.FindAdjacents(test.position))
		})
	}
}

func TestMarkAccessibleRolls(t *testing.T) {
	tests := []struct {
		name     string
		row      RowPosition
		expected string
	}{
		{"mark row 0", 0, "..xx.xx@x."},
		{"mark row 1", 1, "x@@.@.@.@@"},
		{"mark row 2", 2, "@@@@@.x.@@"},
		{"mark row 3", 3, "@.@@@@..@."},
		{"mark row 4", 4, "x@.@@@@.@x"},
		{"mark row 5", 5, ".@@@@@@@.@"},
		{"mark row 6", 6, ".@.@.@.@@@"},
		{"mark row 7", 7, "x.@@@.@@@@"},
		{"mark row 8", 8, ".@@@@@@@@."},
		{"mark row 9", 9, "x.x.@@@.x."},
	}
	grid := NewGrid()
	grid.AddLine("..@@.@@@@.")
	grid.AddLine("@@@.@.@.@@")
	grid.AddLine("@@@@@.@.@@")
	grid.AddLine("@.@@@@..@.")
	grid.AddLine("@@.@@@@.@@")
	grid.AddLine(".@@@@@@@.@")
	grid.AddLine(".@.@.@.@@@")
	grid.AddLine("@.@@@.@@@@")
	grid.AddLine(".@@@@@@@@.")
	grid.AddLine("@.@.@@@.@.")

	grid.MarkAccessibleRolls()

	assert.Equal(t, 13, grid.CountMarkedRolls())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, grid.GetRow(test.row))
		})
	}
}
