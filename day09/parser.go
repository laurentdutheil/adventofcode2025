package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	floor *Floor
}

func NewParser(playground *Floor) *Parser {
	return &Parser{floor: playground}
}

func (p *Parser) ParseLine(line string) {
	coordinates := strings.Split(line, ",")
	c, _ := strconv.Atoi(coordinates[0])
	r, _ := strconv.Atoi(coordinates[1])

	position := Position{Column: c, Row: r}
	p.floor.RedTiles = append(p.floor.RedTiles, position)
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
