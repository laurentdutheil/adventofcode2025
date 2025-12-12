package day09_test

import (
	. "adventofcode2025/day09"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineGreenSegments(t *testing.T) {
	floor := NewFloor()

	floor.RedTiles = append(floor.RedTiles, Position{Column: 7, Row: 1})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 11, Row: 1})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 11, Row: 7})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 9, Row: 7})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 9, Row: 5})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 2, Row: 5})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 2, Row: 3})
	floor.RedTiles = append(floor.RedTiles, Position{Column: 7, Row: 3})

	floor.DetermineGreenSegments()

	assert.Equal(t, &GreenSegment{First: 7, Last: 11}, floor.GetSegment(1))
	assert.Equal(t, &GreenSegment{First: 7, Last: 11}, floor.GetSegment(2))
	assert.Equal(t, &GreenSegment{First: 2, Last: 11}, floor.GetSegment(3))
	assert.Equal(t, &GreenSegment{First: 2, Last: 11}, floor.GetSegment(4))
	assert.Equal(t, &GreenSegment{First: 2, Last: 11}, floor.GetSegment(5))
	assert.Equal(t, &GreenSegment{First: 9, Last: 11}, floor.GetSegment(6))
	assert.Equal(t, &GreenSegment{First: 9, Last: 11}, floor.GetSegment(7))
}
