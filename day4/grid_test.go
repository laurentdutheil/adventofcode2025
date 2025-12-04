package day4_test

import (
	. "adventofcode2025/day4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAdjacents(t *testing.T) {
	tests := []struct {
		name     string
		position []int
		expected string
	}{
		{"adjacents of 5", []int{1, 1}, "12346789"},
		{"adjacents of 1", []int{0, 0}, "....2.45"},
		{"adjacents of 9", []int{2, 2}, "56.8...."},
		{"adjacents of 7", []int{0, 2}, ".45.8..."},
		{"adjacents of 3", []int{2, 0}, "...2.56."},
	}
	grid := NewGrid()
	grid.AddLine("123")
	grid.AddLine("456")
	grid.AddLine("789")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, grid.FindAdjacents(test.position[0], test.position[1]))
		})
	}
}
