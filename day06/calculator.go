package day06

type Calculator struct {
	Numbers   [][]int
	operators []Operator
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

type Operator func(int, int) int

func (o Operator) String() string {
	if o(2, 3) == 5 {
		return "+"
	} else if o(2, 3) == 6 {
		return "*"
	}
	return ""
}

var operators = map[string]Operator{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}
