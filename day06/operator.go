package day06

type Operator struct {
	compute func(int, int) int
	neutral int
	label   string
}

func (o Operator) String() string {
	return o.label
}

func (o Operator) Neutral() int {
	return o.neutral
}

var operators = map[string]Operator{
	"+": {
		compute: func(a, b int) int { return a + b },
		neutral: 0,
		label:   "+",
	},
	"*": {
		compute: func(a, b int) int { return a * b },
		neutral: 1,
		label:   "*",
	},
}
