package day08_test

import (
	. "adventofcode2025/day08"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairDistance(t *testing.T) {
	p1 := NewJunctionBox(11, 12, 13)
	p2 := NewJunctionBox(12, 13, 14)
	pair := NewPair(p1, p2)

	d := pair.Distance()

	assert.Equal(t, math.Sqrt(3), d)
}

func TestPairCombination(t *testing.T) {
	positions := []*JunctionBox{
		{X: 1, Y: 1, Z: 1},
		{X: 2, Y: 2, Z: 2},
		{X: 3, Y: 3, Z: 3},
		{X: 4, Y: 4, Z: 4},
	}

	combinations := PairCombination(positions)

	assert.Equal(t, 6, len(combinations))
	assert.Equal(t, []*Pair{
		NewPair(NewJunctionBox(1, 1, 1), NewJunctionBox(2, 2, 2)),
		NewPair(NewJunctionBox(1, 1, 1), NewJunctionBox(3, 3, 3)),
		NewPair(NewJunctionBox(1, 1, 1), NewJunctionBox(4, 4, 4)),
		NewPair(NewJunctionBox(2, 2, 2), NewJunctionBox(3, 3, 3)),
		NewPair(NewJunctionBox(2, 2, 2), NewJunctionBox(4, 4, 4)),
		NewPair(NewJunctionBox(3, 3, 3), NewJunctionBox(4, 4, 4)),
	}, combinations)
}
