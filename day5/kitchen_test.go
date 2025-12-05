package day5_test

import (
	. "adventofcode2025/day5"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKitchen_CheckFreshness(t *testing.T) {
	tests := []struct {
		name       string
		ingredient *Ingredient
		expected   bool
	}{
		{"1 is spoiled", NewIngredient(1), false},
		{"5 is fresh", NewIngredient(5), true},
		{"8 is spoiled", NewIngredient(8), false},
		{"11 is fresh", NewIngredient(11), true},
		{"17 is fresh", NewIngredient(17), true},
		{"32 is spoiled", NewIngredient(32), false},
	}
	kitchen := Kitchen{}
	kitchen.AddInterval(NewFreshnessInterval(3, 5))
	kitchen.AddInterval(NewFreshnessInterval(10, 14))
	kitchen.AddInterval(NewFreshnessInterval(16, 20))
	kitchen.AddInterval(NewFreshnessInterval(12, 18))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			kitchen.CheckFreshness(test.ingredient)
			assert.Equal(t, test.expected, test.ingredient.Fresh)
		})
	}
}

func TestMergeIntervals(t *testing.T) {
	kitchen := Kitchen{}
	kitchen.AddInterval(NewFreshnessInterval(3, 5))
	kitchen.AddInterval(NewFreshnessInterval(10, 14))
	kitchen.AddInterval(NewFreshnessInterval(16, 20))
	kitchen.AddInterval(NewFreshnessInterval(12, 18))

	kitchen.MergeIntervals()

	expected := []FreshnessInterval{
		NewFreshnessInterval(3, 5),
		NewFreshnessInterval(10, 20),
	}

	assert.Equal(t, expected, kitchen.MergedIntervals)
}

func TestCountFreshIngredients(t *testing.T) {
	kitchen := Kitchen{}
	kitchen.AddInterval(NewFreshnessInterval(3, 5))
	kitchen.AddInterval(NewFreshnessInterval(10, 14))
	kitchen.AddInterval(NewFreshnessInterval(16, 20))
	kitchen.AddInterval(NewFreshnessInterval(12, 18))

	kitchen.MergeIntervals()
	count := kitchen.CountIngredients()

	assert.Equal(t, 14, count)
}
