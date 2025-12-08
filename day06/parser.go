package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseFile() *Calculator {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	calculator := NewCalculator()
	for scanner.Scan() {
		line := scanner.Text()
		ParseLine(line, calculator)
	}
	return calculator
}

func ParseLine(line string, calculator *Calculator) {
	fields := strings.Fields(line)
	for i, field := range fields {
		if field == "*" || field == "+" {
			calculator.AppendOperator(field)
		} else {
			n, _ := strconv.Atoi(field)
			calculator.AddNumberInColumn(n, i)
		}
	}
}
