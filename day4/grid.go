package day4

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

func (g *Grid) getChar(column int, line int) string {
	if line < 0 || column < 0 || line >= len(g.lines) || column >= len(g.lines[line]) {
		return "."
	}
	return string(g.lines[line][column])
}
