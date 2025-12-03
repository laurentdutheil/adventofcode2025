package day2

import (
	"strconv"
)

var primes = []int{2, 3, 5, 7, 11}

func Validate(id int) bool {
	sId := strconv.Itoa(id)
	if len(sId) == 1 {
		return true
	}

	for _, div := range primes {
		i := len(sId) / div
		result := len(sId)%div == 0
		for j := 0; j < div-1; j++ {
			result = result && sId[j*i:(j+1)*i] == sId[(j+1)*i:(j+2)*i]
		}
		if result {
			return false
		}
	}

	return true
}

func ValidateInterval(interval []int) []int {
	var result = make([]int, 0)
	for i := interval[0]; i <= interval[1]; i++ {
		if !Validate(i) {
			result = append(result, i)
		}
	}
	return result
}
