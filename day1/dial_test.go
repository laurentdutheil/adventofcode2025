package day1_test

import (
	"testing"

	. "adventofcode2025/day1"

	"github.com/stretchr/testify/assert"
)

func TestDialStartAt50(t *testing.T) {
	dial := NewDial()
	assert.Equal(t, 50, dial.Position)
}
