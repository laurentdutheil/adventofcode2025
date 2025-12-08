package day06_test

import (
	. "adventofcode2025/day06"

	"github.com/stretchr/testify/assert"
)
import "testing"

func TestCompute(t *testing.T) {
	calculator := NewCalculator()
	ParseLine("123 328  51 64 ", calculator)
	ParseLine(" 45 64  387 23 ", calculator)
	ParseLine("  6 98  215 314", calculator)
	ParseLine("*   +   *   +  ", calculator)

	total := calculator.Compute()
	assert.Equal(t, 4277556, total)
}
