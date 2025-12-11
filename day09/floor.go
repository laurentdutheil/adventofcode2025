package day09

type Floor struct {
	RedTiles []Position
}

func (f Floor) LargestRectangle() int {
	maxArea := 0
	for _, p1 := range f.RedTiles {
		for _, p2 := range f.RedTiles {
			area := Area(p1, p2)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
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
