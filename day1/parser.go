package day1

import "strconv"

type Parser struct {
	dial *Dial
}

func (p *Parser) Parse(line string) {
	direction := line[0]
	clicks, _ := strconv.Atoi(line[1:])
	switch direction {
	case 'L':
		p.dial.TurnLeft(clicks)
	case 'R':
		p.dial.TurnRight(clicks)
	}
}

func NewParser(d *Dial) *Parser {
	return &Parser{dial: d}
}
