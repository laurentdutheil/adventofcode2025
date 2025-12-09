package day07_test

import (
	. "adventofcode2025/day07"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	teleporter := NewTeleporter()
	parser := NewParser(teleporter)

	parser.ParseLine(".......S.......")
	parser.ParseLine("...............")
	parser.ParseLine(".......^.......")
	parser.ParseLine("...............")
	parser.ParseLine("......^.^......")
	parser.ParseLine("...............")
	parser.ParseLine(".....^.^.^.....")
	parser.ParseLine("...............")
	parser.ParseLine("....^.^...^....")
	parser.ParseLine("...............")
	parser.ParseLine("...^.^...^.^...")
	parser.ParseLine("...............")
	parser.ParseLine("..^...^.....^..")
	parser.ParseLine("...............")
	parser.ParseLine(".^.^.^.^.^...^.")
	parser.ParseLine("...............")

	parser.Start()

	assert.Equal(t, 21, teleporter.Splits())
	assert.Equal(t, 40, teleporter.TachyonCount())
}

func TestParseFile(t *testing.T) {
	teleporter := NewTeleporter()
	parser := NewParser(teleporter)

	parser.ParseFile()

	parser.Start()

	assert.Equal(t, 1587, teleporter.Splits())
	assert.Equal(t, 5748679033029, teleporter.TachyonCount())
}
