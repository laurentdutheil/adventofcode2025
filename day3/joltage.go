package day3

import "strconv"

func FindMaxDigit(bank string) (int, int) {
	m := 0
	maxIndex := 0
	for index, c := range bank {
		digit, _ := strconv.Atoi(string(c))
		if digit > m {
			m = digit
			maxIndex = index
		}
	}
	return maxIndex, m
}

func MaxJoltage(text string) int {
	result := 0
	index := 0
	for i := 11; i >= 0; i-- {
		j, digit := FindMaxDigit(text[index : len(text)-i])
		index = index + j + 1
		result = result*10 + digit
	}
	return result
}
