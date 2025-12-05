package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseFile() *Kitchen {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	kitchen := NewKitchen()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			interval := ParseFreshnessInterval(line)
			kitchen.AddInterval(interval)
		} else if line != "" {
			ingredient := ParseIngredient(line)
			kitchen.CheckFreshness(ingredient)
		}
	}
	return kitchen
}

func ParseIngredient(line string) *Ingredient {
	i, _ := strconv.Atoi(line)
	return NewIngredient(i)
}

func ParseFreshnessInterval(line string) FreshnessInterval {
	interval := strings.Split(line, "-")
	mn, _ := strconv.Atoi(interval[0])
	mx, _ := strconv.Atoi(interval[1])
	return NewFreshnessInterval(mn, mx)
}
