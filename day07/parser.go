package day07

import (
	"bufio"
	"os"
)

type Parser struct {
	lines      []string
	teleporter *Teleporter
}

func NewParser(teleporter *Teleporter) *Parser {
	return &Parser{teleporter: teleporter}
}

func (p *Parser) ParseLine(line string) {
	p.lines = append(p.lines, line)
}

func (p *Parser) ParseFile() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		p.ParseLine(line)
	}
}

func (p *Parser) Start() {
	p.teleporter.FirstLine(p.lines[0])
	for i := 1; i < len(p.lines); i++ {
		p.teleporter.Line(p.lines[i])
	}
}
