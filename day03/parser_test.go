package day03_test

import (
	. "adventofcode2025/day03"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	total := ParseFile()
	assert.Equal(t, 171989894144198, total)
}
