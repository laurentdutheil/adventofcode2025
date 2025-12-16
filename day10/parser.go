package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	factory *Factory
}

func NewParser(factory *Factory) *Parser {
	return &Parser{factory: factory}
}

func (p *Parser) ParseLine(line string) {
	firstIndexIndicator := strings.Index(line, "[") + 1
	lastIndexIndicator := strings.Index(line, "]")
	indicator := line[firstIndexIndicator:lastIndexIndicator]

	m := NewMachine(indicator)

	p.parseButtons(line, m)

	p.factory.AddMachin(m)
}

func (p *Parser) parseButtons(line string, m *Machine) {
	firstIndexButtons := strings.Index(line, "(") + 1
	lastIndexButtons := strings.LastIndex(line, ")")
	subLine := line[firstIndexButtons:lastIndexButtons]
	subLine = strings.ReplaceAll(subLine, "(", "")
	subLine = strings.ReplaceAll(subLine, ")", "")
	buttons := strings.Split(subLine, " ")

	for _, b := range buttons {
		lights := strings.Split(b, ",")
		var lightsPosition []int
		for _, l := range lights {
			light, _ := strconv.Atoi(l)
			lightsPosition = append(lightsPosition, light)
		}
		m.AddButton(NewButton(lightsPosition))
	}
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
