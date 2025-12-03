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
		{"987654321111111 max joltage is 987654321111", "987654321111111", 987654321111},
		{"811111111111119 max joltage is 811111111119", "811111111111119", 811111111119},
		{"234234234234278 max joltage is 434234234278", "234234234234278", 434234234278},
		{"818181911112111 max joltage is 888911112111", "818181911112111", 888911112111},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, MaxJoltage(test.bank))
		})
	}
}
