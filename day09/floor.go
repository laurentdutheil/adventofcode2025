package day09

type Floor struct {
	RedTiles      []Position
	GreenSegments map[int]*GreenSegment
}

func NewFloor() *Floor {
	f := &Floor{}
	f.GreenSegments = make(map[int]*GreenSegment)
	return f
}

func (f *Floor) LargestRectangle() int {
	maxArea := 0
	for i := 0; i < len(f.RedTiles); i++ {
		p1 := f.RedTiles[i]
		for j := i; j < len(f.RedTiles); j++ {
			p2 := f.RedTiles[j]
			area := Area(p1, p2)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func (f *Floor) LargestRectanglePart2() int {
	maxArea := 0
	for i := 0; i < len(f.RedTiles); i++ {
		p1 := f.RedTiles[i]
		for j := i; j < len(f.RedTiles); j++ {
			p2 := f.RedTiles[j]
			area := Area(p1, p2)
			if area > maxArea && f.isIncluded(p1, p2) {
				maxArea = area
			}
		}
	}
	return maxArea
}

func (f *Floor) GetSegment(row int) *GreenSegment {
	return f.GreenSegments[row]
}

func (f *Floor) DetermineGreenSegments() {
	var previousRedTile *Position
	for _, tile := range f.RedTiles {
		if previousRedTile != nil {
			if previousRedTile.Row == tile.Row {
				f.updateGreenSegment(tile, previousRedTile)
			} else {
				for row := min(previousRedTile.Row, tile.Row); row <= max(previousRedTile.Row, tile.Row); row++ {
					p := Position{
						Column: tile.Column,
						Row:    row,
					}
					f.updateGreenSegment(p, previousRedTile)
				}
			}
		}
		previousRedTile = &tile
	}

	// link last to first
	tile := f.RedTiles[0]
	if previousRedTile != nil {
		if previousRedTile.Row == tile.Row {
			f.updateGreenSegment(tile, previousRedTile)
		} else {
			for row := min(previousRedTile.Row, tile.Row); row <= max(previousRedTile.Row, tile.Row); row++ {
				p := Position{
					Column: tile.Column,
					Row:    row,
				}
				f.updateGreenSegment(p, previousRedTile)
			}
		}
	}

}

func (f *Floor) updateGreenSegment(tile Position, previousRedTile *Position) {
	_, ok := f.GreenSegments[tile.Row]
	if !ok {
		f.GreenSegments[tile.Row] = &GreenSegment{
			First: min(previousRedTile.Column, tile.Column),
			Last:  max(previousRedTile.Column, tile.Column),
		}
	}
	if f.GreenSegments[tile.Row].First > tile.Column {
		f.GreenSegments[tile.Row].First = tile.Column
	}
	if f.GreenSegments[tile.Row].Last < tile.Column {
		f.GreenSegments[tile.Row].Last = tile.Column
	}
}

func (f *Floor) isIncluded(p1 Position, p2 Position) bool {
	pMin := Position{
		Column: min(p1.Column, p2.Column),
		Row:    min(p1.Row, p2.Row),
	}
	pMax := Position{
		Column: max(p1.Column, p2.Column),
		Row:    max(p1.Row, p2.Row),
	}

	for r := pMin.Row; r <= pMax.Row; r++ {
		if pMin.Column < f.GreenSegments[r].First || pMax.Column > f.GreenSegments[r].Last {
			return false
		}
	}
	return true
}

func Area(p1 Position, p2 Position) int {
	return (1 + Abs(p1.Column-p2.Column)) * (1 + Abs(p1.Row-p2.Row))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Position struct {
	Column int
	Row    int
}

type GreenSegment struct {
	First int
	Last  int
}
