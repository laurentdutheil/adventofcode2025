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

func TestParseLine2(t *testing.T) {
	calculator := NewCalculator()
	ParseLine2("123 328  51 64 ", calculator)
	ParseLine2(" 45 64  387 23 ", calculator)
	ParseLine2("  6 98  215 314", calculator)
	ParseLine2("*   +   *   +  ", calculator)

	calculator.LeftToRightCompute()

	assert.Equal(t, []int{1, 24, 356}, calculator.Column(0))
	assert.Equal(t, []int{369, 248, 8}, calculator.Column(1))
	assert.Equal(t, []int{32, 581, 175}, calculator.Column(2))
	assert.Equal(t, []int{623, 431, 4}, calculator.Column(3))
	assert.Equal(t, []string{"*", "+", "*", "+"}, calculator.Operators())

}

func TestPivot(t *testing.T) {
	lines := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
	}

	p := Pivot(lines)

	assert.Equal(t, []string{
		"1  ", "24 ", "356", "   ",
		"369", "248", "8  ", "   ",
		" 32", "581", "175", "   ",
		"623", "431", "  4",
	}, p)
}

func TestParseFile2(t *testing.T) {
	calculator := ParseFile2()
	total := calculator.LeftToRightCompute()
	assert.Equal(t, 7903168391557, total)
}
