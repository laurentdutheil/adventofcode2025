package day4_test

import (
	"adventofcode2025/day4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessibleRolls(t *testing.T) {
	grid := day4.ParseFile()
	grid.MarkAccessibleRolls()
	assert.Equal(t, 1508, grid.CountMarkedRolls())
}
