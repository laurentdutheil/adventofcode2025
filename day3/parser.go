package day3

import (
	"bufio"
	"os"
)

func ParseFile() int {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total += MaxJoltage(scanner.Text())
	}
	return total
}
