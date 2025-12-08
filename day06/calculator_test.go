package day06_test

import (
	. "adventofcode2025/day06"

	"github.com/stretchr/testify/assert"
)
import "testing"

func TestCompute(t *testing.T) {
	calculator := NewCalculator()
	//"123 328  51 64 "
	//" 45 64  387 23 "
	//"  6 98  215 314"
	//"*   +   *   +  "
	calculator.AddNumberInColumn(123, 0)
	calculator.AddNumberInColumn(45, 0)
	calculator.AddNumberInColumn(6, 0)
	calculator.AppendOperator("*")
	calculator.AddNumberInColumn(328, 1)
	calculator.AddNumberInColumn(64, 1)
	calculator.AddNumberInColumn(98, 1)
	calculator.AppendOperator("+")
	calculator.AddNumberInColumn(51, 2)
	calculator.AddNumberInColumn(387, 2)
	calculator.AddNumberInColumn(215, 2)
	calculator.AppendOperator("*")
	calculator.AddNumberInColumn(64, 3)
	calculator.AddNumberInColumn(23, 3)
	calculator.AddNumberInColumn(314, 3)
	calculator.AppendOperator("+")

	total := calculator.Compute()
	assert.Equal(t, 4277556, total)
}
