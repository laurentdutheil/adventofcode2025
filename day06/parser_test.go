package day06_test

import (
	. "adventofcode2025/day06"

	"github.com/stretchr/testify/assert"
)
import "testing"

func TestParseLine(t *testing.T) {
	calculator := NewCalculator()
	ParseLine("123 328  51 64 ", calculator)
	ParseLine(" 45 64  387 23 ", calculator)
	ParseLine("  6 98  215 314", calculator)
	ParseLine("*   +   *   +  ", calculator)

	assert.Equal(t, []int{123, 45, 6}, calculator.Column(0))
	assert.Equal(t, []int{328, 64, 98}, calculator.Column(1))
	assert.Equal(t, []int{51, 387, 215}, calculator.Column(2))
	assert.Equal(t, []int{64, 23, 314}, calculator.Column(3))
	assert.Equal(t, []string{"*", "+", "*", "+"}, calculator.Operators())
}

func TestParseFile(t *testing.T) {
	calculator := ParseFile()
	total := calculator.Compute()
	assert.Equal(t, 4076006202939, total)
}
