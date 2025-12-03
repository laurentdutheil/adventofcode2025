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
		{"111 is invalid", 111, false},
		{"1010 is invalid", 1010, false},
		{"222222 is invalid", 222222, false},
		{"565656 is invalid", 565656, false},
		{"1111111 is invalid", 1111111, false},
		{"2121212121 is invalid", 2121212121, false},
		{"55555555555 is invalid", 55555555555, false},
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
		{"11-22 has two invalid IDs", []int{11, 22}, []int{11, 22}},
		{"95-115 has one invalid ID", []int{95, 115}, []int{99, 111}},
		{"998-1012 has one invalid ID", []int{998, 1012}, []int{999, 1010}},
		{"1188511880-1188511890 has one invalid ID", []int{1188511880, 1188511890}, []int{1188511885}},
		{"222220-222224 has one invalid ID", []int{222220, 222224}, []int{222222}},
		{"1698522-1698528 has no invalid ID", []int{1698522, 1698528}, []int{}},
		{"446443-446449 has one invalid ID", []int{446443, 446449}, []int{446446}},
		{"38593856-38593862 has one invalid ID", []int{38593856, 38593862}, []int{38593859}},
		{"565653-565659 now has one invalid ID", []int{565653, 565659}, []int{565656}},
		{"824824821-824824827 now has one invalid ID", []int{824824821, 824824827}, []int{824824824}},
		{"2121212118-2121212124 now has one invalid ID", []int{2121212118, 2121212124}, []int{2121212121}},
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
	assert.Equal(t, 46270373595, result)
}
