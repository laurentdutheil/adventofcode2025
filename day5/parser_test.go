package day5_test

import (
	. "adventofcode2025/day5"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFreshnessInterval(t *testing.T) {
	interval := ParseFreshnessInterval("3-5")
	assert.Equal(t, NewFreshnessInterval(3, 5), interval)
}

func TestPart1(t *testing.T) {
	kitchen := ParseFile()
	assert.Equal(t, 563, kitchen.Count)
}

func TestPart2(t *testing.T) {
	kitchen := ParseFile()
	kitchen.MergeIntervals()
	count := kitchen.CountIngredients()
	assert.Equal(t, 338693411431456, count)
}
