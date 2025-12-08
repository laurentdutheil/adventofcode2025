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
			cumulative = o(n, cumulative)
		}
		result += cumulative
	}
	return result
}

func (c *Calculator) LeftToRightCompute() int {
	pivot := Pivot(c.lines[:len(c.lines)-1])

	i := 0
	for _, s := range pivot {
		if strings.TrimSpace(s) == "" {
			i++
			continue
		}
		n, _ := strconv.Atoi(strings.TrimSpace(s))
		c.AddNumberInColumn(n, i)
	}

	ParseLine(c.lines[len(c.lines)-1], c)

	return c.Compute()
}

func Pivot(lines []string) []string {
	var p []string
	l := len(lines[0])
	for j := 0; j < l; j++ {
		p = append(p, "")
		for _, line := range lines {
			p[j] = p[j] + string(line[j])
		}
	}

	return p
}

type Operator func(int, int) int

func (o Operator) String() string {
	if o(2, 3) == 5 {
		return "+"
	} else if o(2, 3) == 6 {
		return "*"
	}
	return ""
}

func (o Operator) Neutral() int {
	if o.String() == "*" {
		return 1
	}
	return 0
}

var operators = map[string]Operator{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}
