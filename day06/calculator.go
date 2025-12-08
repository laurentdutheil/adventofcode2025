package day06

import (
	"strconv"
	"strings"
)

type Calculator struct {
	Numbers   [][]int
	operators []Operator
	lines     []string
}

func NewCalculator() *Calculator {
	c := &Calculator{}
	c.Numbers = [][]int{}
	return c
}

func (c *Calculator) Column(i int) []int {
	return c.Numbers[i]
}

func (c *Calculator) AddNumberInColumn(n int, column int) {
	if len(c.Numbers) <= column {
		c.Numbers = append(c.Numbers, []int{})
	}
	c.Numbers[column] = append(c.Numbers[column], n)
}

func (c *Calculator) AppendOperator(o string) {
	c.operators = append(c.operators, operators[o])
}

func (c *Calculator) Operators() []string {
	var result []string
	for _, o := range c.operators {
		result = append(result, o.String())
	}
	return result
}

func (c *Calculator) Compute() int {
	result := 0
	for i := 0; i < len(c.Numbers); i++ {
		o := c.operators[i]
		cumulative := o.Neutral()
		for _, n := range c.Numbers[i] {
			cumulative = o.compute(n, cumulative)
		}
		result += cumulative
	}
	return result
}

func (c *Calculator) LeftToRightCompute() int {
	pivot := Pivot(c.lines[:len(c.lines)-1])

	column := 0
	for _, s := range pivot {
		if strings.TrimSpace(s) == "" {
			column++
		} else {
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			c.AddNumberInColumn(n, column)
		}
	}

	ParseLine(c.lines[len(c.lines)-1], c)

	return c.Compute()
}

func Pivot(lines []string) []string {
	var result []string
	for j := 0; j < len(lines[0]); j++ {
		result = append(result, "")
		for _, line := range lines {
			result[j] += string(line[j])
		}
	}
	return result
}
