package day08_test

import (
	. "adventofcode2025/day08"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseLine(t *testing.T) {
	playground := NewPlayground()
	parser := NewParser(playground)

	parser.ParseLine("162,817,812")

	assert.Len(t, playground.JunctionBoxes, 1)
	assert.Contains(t, playground.JunctionBoxes, NewJunctionBox(162, 817, 812))
}

func TestParser_Combination(t *testing.T) {
	playground := NewPlayground()
	parser := NewParser(playground)

	parser.ParseLine("162,817,812")
	parser.ParseLine("57,618,57")
	parser.ParseLine("906,360,560")
	parser.ParseLine("592,479,940")
	parser.ParseLine("352,342,300")
	parser.ParseLine("466,668,158")
	parser.ParseLine("542,29,236")
	parser.ParseLine("431,825,988")
	parser.ParseLine("739,650,466")
	parser.ParseLine("52,470,668")
	parser.ParseLine("216,146,977")
	parser.ParseLine("819,987,18")
	parser.ParseLine("117,168,530")
	parser.ParseLine("805,96,715")
	parser.ParseLine("346,949,466")
	parser.ParseLine("970,615,88")
	parser.ParseLine("941,993,340")
	parser.ParseLine("862,61,35")
	parser.ParseLine("984,92,344")
	parser.ParseLine("425,690,689")

	playground.SortCombinations()

	assert.Equal(t, 190, len(playground.Combinations))
	shortestPair := NewPair(NewJunctionBox(162, 817, 812), NewJunctionBox(425, 690, 689))
	assert.Equal(t, shortestPair, playground.Combinations[0])

	playground.Group(10)

	result := playground.CountPart1()
	assert.Equal(t, 40, result)
}

func TestParser_Combination_Part2(t *testing.T) {
	playground := NewPlayground()
	parser := NewParser(playground)

	parser.ParseLine("162,817,812")
	parser.ParseLine("57,618,57")
	parser.ParseLine("906,360,560")
	parser.ParseLine("592,479,940")
	parser.ParseLine("352,342,300")
	parser.ParseLine("466,668,158")
	parser.ParseLine("542,29,236")
	parser.ParseLine("431,825,988")
	parser.ParseLine("739,650,466")
	parser.ParseLine("52,470,668")
	parser.ParseLine("216,146,977")
	parser.ParseLine("819,987,18")
	parser.ParseLine("117,168,530")
	parser.ParseLine("805,96,715")
	parser.ParseLine("346,949,466")
	parser.ParseLine("970,615,88")
	parser.ParseLine("941,993,340")
	parser.ParseLine("862,61,35")
	parser.ParseLine("984,92,344")
	parser.ParseLine("425,690,689")

	playground.SortCombinations()

	assert.Equal(t, 190, len(playground.Combinations))
	shortestPair := NewPair(NewJunctionBox(162, 817, 812), NewJunctionBox(425, 690, 689))
	assert.Equal(t, shortestPair, playground.Combinations[0])

	last := playground.Group(190)

	result := last.DistanceWall()
	assert.Equal(t, 25272, result)
}

func TestParser_ParseFile(t *testing.T) {
	playground := NewPlayground()
	parser := NewParser(playground)

	parser.ParseFile()

	playground.SortCombinations()
	playground.Group(1000)
	result := playground.CountPart1()

	assert.Equal(t, 131150, result)
}

func TestParser_ParseFile_Part2(t *testing.T) {
	playground := NewPlayground()
	parser := NewParser(playground)

	parser.ParseFile()

	playground.SortCombinations()
	last := playground.Group(len(playground.Combinations))

	result := last.DistanceWall()

	assert.Equal(t, 2497445, result)
}
