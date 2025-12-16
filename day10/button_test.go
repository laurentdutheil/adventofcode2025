package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToggle(t *testing.T) {
	button := NewButton([]int{1, 2})

	indicator := button.Toggle("...")
	assert.Equal(t, ".##", indicator)

	indicator = button.Toggle("##.")
	assert.Equal(t, "#.#", indicator)
}
