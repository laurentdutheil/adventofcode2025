package day08

import "math"

type JunctionBox struct {
	X       int
	Y       int
	Z       int
	Circuit int
}

func NewJunctionBox(x int, y int, z int) *JunctionBox {
	return &JunctionBox{X: x, Y: y, Z: z}
}

type Pair struct {
	p1 *JunctionBox
	p2 *JunctionBox
}

func (pair Pair) Distance() float64 {
	p := pair.p1
	p2 := pair.p2
	dx := math.Pow(float64(p.X-p2.X), 2)
	dy := math.Pow(float64(p.Y-p2.Y), 2)
	dz := math.Pow(float64(p.Z-p2.Z), 2)

	return math.Sqrt(dx + dy + dz)
}

func (pair Pair) DistanceWall() int {
	return pair.p1.X * pair.p2.X
}

func NewPair(p1 *JunctionBox, p2 *JunctionBox) *Pair {
	return &Pair{p1: p1, p2: p2}
}

func PairCombination(junctionBoxes []*JunctionBox) []*Pair {
	var result []*Pair
	for i := 0; i < len(junctionBoxes)-1; i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			result = append(result, NewPair(junctionBoxes[i], junctionBoxes[j]))
		}
	}
	return result
}
