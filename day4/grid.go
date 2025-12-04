package day4

import "strings"

type Grid struct {
	lines []string
}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) AddLine(line string) {
	g.lines = append(g.lines, line)
}

func (g *Grid) FindAdjacents(column int, line int) string {
	result := ""
	for l := line - 1; l <= line+1; l++ {
		for c := column - 1; c <= column+1; c++ {
			if l != line || c != column {
				result += g.getChar(c, l)
			}
		}
	}
	return result
}

func (g *Grid) CountMarkedRolls() int {
	result := 0
	for _, line := range g.lines {
		result += strings.Count(line, "x")
	}
	return result
}

func (g *Grid) GetLine(line int) string {
	return g.lines[line]
}

func (g *Grid) MarkAccessibleRolls() {
	for l := 0; l < len(g.lines); l++ {
		g.markLine(l)
	}
}

func (g *Grid) markLine(line int) {
	for c := 0; c < len(g.lines[line]); c++ {
		adjacents := g.FindAdjacents(c, line)
		adjacents = strings.ReplaceAll(adjacents, ".", "")
		if g.getChar(c, line) != "." && len(adjacents) < 4 {
			g.lines[line] = replaceAtIndex(g.lines[line], 'x', c)
		}
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func (g *Grid) getChar(column int, line int) string {
	if line < 0 || column < 0 || line >= len(g.lines) || column >= len(g.lines[line]) {
		return "."
	}
	return string(g.lines[line][column])
}
