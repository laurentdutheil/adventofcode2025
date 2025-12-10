package day08

import (
	"maps"
	"slices"
	"sort"
)

type Playground struct {
	JunctionBoxes []*JunctionBox
	Combinations  []*Pair
}

func NewPlayground() *Playground {
	return &Playground{}
}

func (p *Playground) SortCombinations() {
	p.Combinations = PairCombination(p.JunctionBoxes)
	slices.SortFunc(p.Combinations, func(p1, p2 *Pair) int {
		if p1.Distance() < p2.Distance() {
			return -1
		}
		return 1
	})
}

func (p *Playground) Group(nbCables int) {
	currentCircuit := 0
	for i := 0; i < nbCables; i++ {
		pair := p.Combinations[i]
		if pair.p1.Circuit == 0 && pair.p2.Circuit == 0 {
			currentCircuit++
			pair.p1.Circuit = currentCircuit
			pair.p2.Circuit = currentCircuit
		} else if pair.p1.Circuit != 0 && pair.p2.Circuit == 0 {
			pair.p2.Circuit = pair.p1.Circuit
		} else if pair.p2.Circuit != 0 && pair.p1.Circuit == 0 {
			pair.p1.Circuit = pair.p2.Circuit
		} else if pair.p1.Circuit != pair.p2.Circuit {
			p.mergeCircuits(pair.p1.Circuit, pair.p2.Circuit)
		}
	}
}

func (p *Playground) mergeCircuits(c1 int, c2 int) {
	for _, box := range p.JunctionBoxes {
		if box.Circuit == c2 {
			box.Circuit = c1
		}
	}
}

func (p *Playground) CountPart1() int {
	var counters = make(map[int]int)

	for _, box := range p.JunctionBoxes {
		if box.Circuit != 0 {
			counters[box.Circuit]++
		}
	}

	v := slices.Collect(maps.Values(counters))
	sort.Sort(sort.Reverse(sort.IntSlice(v)))

	return v[0] * v[1] * v[2]
}
