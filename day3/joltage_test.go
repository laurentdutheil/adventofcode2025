package day3_test

import (
	"testing"

	. "adventofcode2025/day3"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxDigit(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		expected int
	}{
		{"987654321111111 max digit", "987654321111111", 9},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, digit := FindMaxDigit(test.bank)
			assert.Equal(t, test.expected, digit)
		})
	}
}

func TestFindMaxJoltage(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		expected int
	}{
		{"987654321111111 max joltage is 98", "987654321111111", 98},
		{"811111111111119 max joltage is 89", "811111111111119", 89},
		{"234234234234278 max joltage is 78", "234234234234278", 78},
		{"818181911112111 max joltage is 92", "818181911112111", 92},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, MaxJoltage(test.bank))
		})
	}
}
