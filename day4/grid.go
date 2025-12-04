package day4

import "strings"

type Grid struct {
	rows []string
}

type ColumnPosition int
type RowPosition int

type Position struct {
	Column ColumnPosition
	Row    RowPosition
}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) AddLine(line string) {
	g.rows = append(g.rows, line)
}

func (g *Grid) FindAdjacents(position Position) string {
	result := ""
	row := position.Row
	column := position.Column
	for l := row - 1; l <= row+1; l++ {
		for c := column - 1; c <= column+1; c++ {
			if l != row || c != column {
				result += g.getChar(Position{c, l})
			}
		}
	}
	return result
}

func (g *Grid) CountMarkedRolls() int {
	result := 0
	for _, line := range g.rows {
		result += strings.Count(line, "x")
	}
	return result
}

func (g *Grid) GetRow(row RowPosition) string {
	return g.rows[row]
}

func (g *Grid) MarkAccessibleRolls() {
	for r := RowPosition(0); int(r) < len(g.rows); r++ {
		g.markRow(r)
	}
}

func (g *Grid) RemoveAccessibleRolls() {
	for i, line := range g.rows {
		g.rows[i] = strings.ReplaceAll(line, "x", ".")
	}
}

func (g *Grid) markRow(row RowPosition) {
	for c := ColumnPosition(0); int(c) < len(g.rows[row]); c++ {
		if g.getChar(Position{c, row}) != "." {
			adjacents := g.FindAdjacents(Position{c, row})
			adjacents = strings.ReplaceAll(adjacents, ".", "")
			if len(adjacents) < 4 {
				g.rows[row] = replaceAtIndex(g.rows[row], 'x', c)
			}
		}
	}
}

func replaceAtIndex(in string, r rune, row ColumnPosition) string {
	out := []rune(in)
	out[row] = r
	return string(out)
}

func (g *Grid) getChar(position Position) string {
	row := position.Row
	column := position.Column
	if row < 0 || column < 0 || int(row) >= len(g.rows) || int(column) >= len(g.rows[row]) {
		return "."
	}
	return string(g.rows[row][column])
}
