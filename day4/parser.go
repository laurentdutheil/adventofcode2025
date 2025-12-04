package day4

import (
	"bufio"
	"os"
)

func ParseFile() *Grid {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := NewGrid()
	for scanner.Scan() {
		grid.AddLine(scanner.Text())
	}
	return grid
}
