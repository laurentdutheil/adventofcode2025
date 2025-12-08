package day02

import (
	"strconv"
)

var primes = []int{2, 3, 5, 7, 11}

func Validate(id int) bool {
	sId := strconv.Itoa(id)

	for _, nbPart := range primes {
		if len(sId)%nbPart != 0 {
			continue
		}
		if allPartsAreEqual(sId, nbPart) {
			return false
		}
	}

	return true
}

func allPartsAreEqual(sId string, div int) bool {
	nbChar := len(sId) / div
	result := true
	for i := 0; i < div-1; i++ {
		result = result && sId[i*nbChar:(i+1)*nbChar] == sId[(i+1)*nbChar:(i+2)*nbChar]
	}
	return result
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
