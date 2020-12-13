package util

import "strconv"

// ToNumber ...
func ToNumber(value string) int {
	if value == "" {
		return 0
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return num
}
