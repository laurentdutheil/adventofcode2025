package day4

import (
	"iter"
	"strings"
)

type Grid struct {
	rows []string
}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) AddLine(line string) {
	g.rows = append(g.rows, line)
}

func (g *Grid) MarkAccessibleRolls() {
	for p, adjacents := range g.AllAdjacents() {
		if g.getChar(p) != "." {
			rolls := strings.ReplaceAll(adjacents, ".", "")
			if len(rolls) < 4 {
				g.rows[p.Row] = replaceAtIndex(g.rows[p.Row], 'x', p.Column)
			}
		}
	}
}

func (g *Grid) CountMarkedRolls() int {
	result := 0
	for _, line := range g.rows {
		result += strings.Count(line, "x")
	}
	return result
}

func (g *Grid) RemoveAccessibleRolls() {
	for i, line := range g.rows {
		g.rows[i] = strings.ReplaceAll(line, "x", ".")
	}
}

func (g *Grid) All() iter.Seq2[Position, string] {
	return func(yield func(Position, string) bool) {
		for r, row := range g.rows {
			for c, elem := range row {
				if !yield(Position{ColumnPosition(c), RowPosition(r)}, string(elem)) {
					return
				}
			}
		}
	}
}

func (g *Grid) AllRows() iter.Seq2[RowPosition, string] {
	return func(yield func(RowPosition, string) bool) {
		for r, row := range g.rows {
			if !yield(RowPosition(r), row) {
				return
			}
		}
	}
}

func (g *Grid) AllAdjacents() iter.Seq2[Position, string] {
	return func(yield func(Position, string) bool) {
		for position := range g.All() {
			adjacents := g.findAdjacents(position)
			if !yield(position, adjacents) {
				return
			}
		}
	}
}

func (g *Grid) findAdjacents(position Position) string {
	result := ""
	for _, p := range position.Adjacents() {
		result += g.getChar(p)
	}
	return result
}

func (g *Grid) getChar(position Position) string {
	if g.isInGrid(position) {
		return string(g.rows[position.Row][position.Column])
	}
	return ""
}

func (g *Grid) isInGrid(position Position) bool {
	return 0 <= position.Row && int(position.Row) < len(g.rows) && 0 <= position.Column && int(position.Column) < len(g.rows[0])
}

func replaceAtIndex(in string, r rune, row ColumnPosition) string {
	out := []rune(in)
	out[row] = r
	return string(out)
}
