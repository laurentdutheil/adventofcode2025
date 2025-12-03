package day2_test

import (
	"testing"

	. "adventofcode2025/day2"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name   string
		id     int
		expect bool
	}{
		{"Odd length IDs are valid", 3, true},
		{"11 is invalid", 11, false},
		{"12 is valid", 12, true},
		{"1010 is invalid", 1010, false},
		{"222222 is invalid", 222222, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Validate(test.id))
		})
	}
}

func TestValidateInterval(t *testing.T) {
	tests := []struct {
		name     string
		interval []int
		expected []int
	}{
		{"95-115 has one invalid ID", []int{95, 115}, []int{99}},
		{"11-22 has two invalid IDs", []int{11, 22}, []int{11, 22}},
		{"1188511880-1188511890 has one invalid ID", []int{1188511880, 1188511890}, []int{1188511885}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ValidateInterval(test.interval))
		})
	}
}

func TestFile(t *testing.T) {
	result := 0
	intervals := ParseFile()
	for _, sInterval := range intervals {
		interval := ParseInterval(sInterval)
		invalidIds := ValidateInterval(interval)
		for _, invalidId := range invalidIds {
			result += invalidId
		}
	}
	assert.Equal(t, 30599400849, result)
}
