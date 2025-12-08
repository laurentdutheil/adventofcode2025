package day03

import (
	"math"
	"strconv"
)

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

func MaxJoltage(bank string) int {
	return maxJoltageRecursive(bank, 12)
}

func maxJoltageRecursive(bank string, nbDigit int) int {
	if nbDigit == 1 {
		_, digit := FindMaxDigit(bank)
		return digit
	}
	nbDigit--
	i, digit := FindMaxDigit(bank[:len(bank)-nbDigit])
	return digit*int(math.Pow10(nbDigit)) + maxJoltageRecursive(bank[i+1:], nbDigit)
}
