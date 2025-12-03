package day2

import "strconv"

func Validate(id int) bool {
	sId := strconv.Itoa(id)
	i := len(sId) / 2
	return !(sId[:i] == sId[i:])
}

func ValidateInterval(interval []int) []int {
	var result []int
	for i := interval[0]; i <= interval[1]; i++ {
		if !Validate(i) {
			result = append(result, i)
		}
	}
	return result
}
