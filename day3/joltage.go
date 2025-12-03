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
	index, tens := FindMaxDigit(text[:len(text)-1])
	_, units := FindMaxDigit(text[index+1:])
	return tens*10 + units
}
