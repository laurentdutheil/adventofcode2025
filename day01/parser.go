package day01

import (
	"bufio"
	"os"
	"strconv"
)

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

func (p *Parser) ParseFile(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.Parse(scanner.Text())
	}
}

func NewParser(d *Dial) *Parser {
	return &Parser{dial: d}
}
