package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseLine(line string) []string {
	return strings.Split(line, ",")
}

func ParseInterval(interval string) []int {
	sId := strings.Split(interval, "-")
	result := make([]int, 2)
	result[0], _ = strconv.Atoi(sId[0])
	result[1], _ = strconv.Atoi(sId[1])
	return result
}

func ParseFile() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return ParseLine(scanner.Text())
	}
	return nil
}
