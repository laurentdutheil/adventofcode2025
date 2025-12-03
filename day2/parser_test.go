package day2_test

import (
	"testing"

	. "adventofcode2025/day2"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected []string
	}{
		{"one interval", "11-22", []string{"11-22"}},
		{"two interval", "11-22,95-115", []string{"11-22", "95-115"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ParseLine(test.line))
		})
	}
}

func TestParseInterval(t *testing.T) {
	tests := []struct {
		name     string
		interval string
		expected []int
	}{
		{"one interval", "11-22", []int{11, 22}},
		{"two interval", "95-115", []int{95, 115}},
		{"two interval", "2121212118-2121212124", []int{2121212118, 2121212124}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ParseInterval(test.interval))
		})
	}
}
