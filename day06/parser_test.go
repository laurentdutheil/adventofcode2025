package day06_test

import (
	. "adventofcode2025/day06"

	"github.com/stretchr/testify/assert"
)
import "testing"

func TestParseLineLeftToRight(t *testing.T) {
	calculator := NewCalculator()
	p := NewParser(calculator)
	p.ParseLine("123 328  51 64 ")
	p.ParseLine(" 45 64  387 23 ")
	p.ParseLine("  6 98  215 314")
	p.ParseLine("*   +   *   +  ")

	p.LeftToRightParse()
	p.OperatorParse()

	assert.Equal(t, []int{123, 45, 6}, calculator.Column(0))
	assert.Equal(t, []int{328, 64, 98}, calculator.Column(1))
	assert.Equal(t, []int{51, 387, 215}, calculator.Column(2))
	assert.Equal(t, []int{64, 23, 314}, calculator.Column(3))
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

func TestParseRightToLeft(t *testing.T) {
	calculator := NewCalculator()
	p := NewParser(calculator)
	p.ParseLine("123 328  51 64 ")
	p.ParseLine(" 45 64  387 23 ")
	p.ParseLine("  6 98  215 314")
	p.ParseLine("*   +   *   +  ")

	p.RightToLeftParse()
	p.OperatorParse()

	assert.Equal(t, []int{1, 24, 356}, calculator.Column(0))
	assert.Equal(t, []int{369, 248, 8}, calculator.Column(1))
	assert.Equal(t, []int{32, 581, 175}, calculator.Column(2))
	assert.Equal(t, []int{623, 431, 4}, calculator.Column(3))
	assert.Equal(t, []string{"*", "+", "*", "+"}, calculator.Operators())
}

func TestParseFilePart1(t *testing.T) {
	calculator := NewCalculator()
	p := NewParser(calculator)
	p.ParseFile()
	p.LeftToRightParse()
	p.OperatorParse()
	total := calculator.Compute()
	assert.Equal(t, 4076006202939, total)
}

func TestParseFilePart2(t *testing.T) {
	calculator := NewCalculator()
	p := NewParser(calculator)
	p.ParseFile()
	p.RightToLeftParse()
	p.OperatorParse()
	total := calculator.Compute()
	assert.Equal(t, 7903168391557, total)
}
