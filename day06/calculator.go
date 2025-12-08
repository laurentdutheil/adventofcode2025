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
