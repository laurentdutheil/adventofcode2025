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
		{"adjacents of 1", Position{Column: 0, Row: 0}, "245"},
		{"adjacents of 2", Position{Column: 1, Row: 0}, "13456"},
		{"adjacents of 3", Position{Column: 2, Row: 0}, "256"},
		{"adjacents of 4", Position{Column: 0, Row: 1}, "12578"},
		{"adjacents of 5", Position{Column: 1, Row: 1}, "12346789"},
		{"adjacents of 6", Position{Column: 2, Row: 1}, "23589"},
		{"adjacents of 7", Position{Column: 0, Row: 2}, "458"},
		{"adjacents of 8", Position{Column: 1, Row: 2}, "45679"},
		{"adjacents of 9", Position{Column: 2, Row: 2}, "568"},
	}
	grid := NewGrid()
	grid.AddLine("123")
	grid.AddLine("456")
	grid.AddLine("789")

	for position, adjacents := range grid.AllAdjacents() {
		test := tests[3*int(position.Row)+int(position.Column)]
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, adjacents)
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

	for i, r := range grid.AllRows() {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, r)
		})
	}
}
