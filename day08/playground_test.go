package day08_test

import (
	. "adventofcode2025/day08"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayground_SortCombinations(t *testing.T) {
	playground := NewPlayground()

	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(1, 1, 1))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(5, 5, 5))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(2, 2, 2))

	playground.SortCombinations()

	assert.Equal(t, []*Pair{
		NewPair(NewJunctionBox(1, 1, 1), NewJunctionBox(2, 2, 2)),
		NewPair(NewJunctionBox(5, 5, 5), NewJunctionBox(2, 2, 2)),
		NewPair(NewJunctionBox(1, 1, 1), NewJunctionBox(5, 5, 5)),
	}, playground.Combinations)
}

func TestGroup0JunctionBoxes(t *testing.T) {
	playground := NewPlayground()

	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(1, 1, 1))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(5, 5, 5))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(2, 2, 2))

	playground.SortCombinations()
	playground.Group(0)

	assert.Equal(t, 0, playground.JunctionBoxes[0].Circuit)
	assert.Equal(t, 0, playground.JunctionBoxes[1].Circuit)
	assert.Equal(t, 0, playground.JunctionBoxes[2].Circuit)
}

func TestGroup1JunctionBoxes(t *testing.T) {
	playground := NewPlayground()

	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(1, 1, 1))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(5, 5, 5))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(2, 2, 2))

	playground.SortCombinations()
	playground.Group(1)

	assert.Equal(t, 1, playground.JunctionBoxes[0].Circuit)
	assert.Equal(t, 0, playground.JunctionBoxes[1].Circuit)
	assert.Equal(t, 1, playground.JunctionBoxes[2].Circuit)
}

func TestGroup3JunctionBoxes(t *testing.T) {
	playground := NewPlayground()

	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(1, 1, 1))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(5, 5, 5))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(2, 2, 2))

	playground.SortCombinations()
	playground.Group(3)

	assert.Equal(t, 1, playground.JunctionBoxes[0].Circuit)
	assert.Equal(t, 1, playground.JunctionBoxes[1].Circuit)
	assert.Equal(t, 1, playground.JunctionBoxes[2].Circuit)
}

func TestGroupJunctionBoxes2Circuits(t *testing.T) {
	playground := NewPlayground()

	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(162, 817, 812))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(425, 690, 689))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(431, 825, 988))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(906, 360, 560))
	playground.JunctionBoxes = append(playground.JunctionBoxes, NewJunctionBox(805, 96, 715))

	playground.SortCombinations()
	playground.Group(4)

	assert.Equal(t, 1, playground.JunctionBoxes[0].Circuit)
	assert.Equal(t, 1, playground.JunctionBoxes[1].Circuit)
	assert.Equal(t, 1, playground.JunctionBoxes[2].Circuit)
	assert.Equal(t, 2, playground.JunctionBoxes[3].Circuit)
	assert.Equal(t, 2, playground.JunctionBoxes[4].Circuit)
}
