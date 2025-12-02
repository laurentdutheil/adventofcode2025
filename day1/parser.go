package day1

import "strconv"

type Parser struct {
	dial *Dial
}

func (p *Parser) Parse(line string) {
	clicks, _ := strconv.Atoi(line[1:])
	p.dial.TurnRight(clicks)
}

func NewParser(d *Dial) *Parser {
	return &Parser{dial: d}
}
