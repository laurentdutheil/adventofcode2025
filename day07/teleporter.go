package day07

import (
	"strings"
)

type Teleporter struct {
	tachyons []int
	splits   int
}

func NewTeleporter() *Teleporter {
	return &Teleporter{}
}

func (t *Teleporter) TachyonCount() int {
	result := 0
	for _, tachyon := range t.tachyons {
		result += tachyon
	}
	return result
}

func (t *Teleporter) Splits() int {
	return t.splits
}

func (t *Teleporter) FirstLine(s string) {
	t.tachyons = make([]int, len(s))
	i := strings.Index(s, "S")
	t.tachyons[i] = 1
}

func (t *Teleporter) Line(s string) {
	temp := make([]int, len(t.tachyons))
	for i := 0; i < len(s); i++ {
		if t.tachyons[i] > 0 {
			if s[i] == '^' {
				temp[i-1] += t.tachyons[i]
				temp[i+1] += t.tachyons[i]
				t.splits++
			} else {
				temp[i] += t.tachyons[i]
			}
		}
	}
	t.tachyons = temp
}
