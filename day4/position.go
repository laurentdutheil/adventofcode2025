package day4

type ColumnPosition int
type RowPosition int

type Position struct {
	Column ColumnPosition
	Row    RowPosition
}

func (p Position) Adjacents() []Position {
	result := make([]Position, 0)
	row := p.Row
	column := p.Column
	for r := row - 1; r <= row+1; r++ {
		for c := column - 1; c <= column+1; c++ {
			if r != row || c != column {
				result = append(result, Position{c, r})
			}
		}
	}
	return result
}
