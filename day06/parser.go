package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	lines      []string
	calculator *Calculator
}

func NewParser(calculator *Calculator) *Parser {
	return &Parser{calculator: calculator}
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

func (p *Parser) LeftToRightParse() {
	for _, line := range p.lines[:len(p.lines)-1] {
		fields := strings.Fields(line)
		for i, field := range fields {
			n, _ := strconv.Atoi(field)
			p.calculator.AddNumberInColumn(n, i)
		}
	}
}

func (p *Parser) RightToLeftParse() {
	pivot := Pivot(p.lines[:len(p.lines)-1])

	column := 0
	for _, s := range pivot {
		if strings.TrimSpace(s) == "" {
			column++
		} else {
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			p.calculator.AddNumberInColumn(n, column)
		}
	}
}

func (p *Parser) OperatorParse() {
	fields := strings.Fields(p.lines[len(p.lines)-1])
	for _, field := range fields {
		p.calculator.AppendOperator(field)
	}

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
