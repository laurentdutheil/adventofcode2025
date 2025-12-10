package day08

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	playground *Playground
}

func NewParser(playground *Playground) *Parser {
	return &Parser{playground: playground}
}

func (p *Parser) ParseLine(line string) {
	coordinates := strings.Split(line, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	z, _ := strconv.Atoi(coordinates[2])
	p.playground.JunctionBoxes = append(p.playground.JunctionBoxes, NewJunctionBox(x, y, z))
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
