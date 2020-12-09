package part2

import (
	"regexp"

	"github.com/barthogenes/adventofcode2020/day2"
)

// GetAnswer Get the answer for the Day 2 puzzle.
func GetAnswer(input []day2.PasswordPolicySet) int {
	count := 0
	for _, element := range input {
		if !isValid(element) {
			continue
		}

		count++
	}

	return count
}

func isValid(input day2.PasswordPolicySet) bool {
	re := regexp.MustCompile(input.Char)
	found := re.FindAllStringIndex(input.Password, -1)
	isValid := false
	for _, element := range found {
		index := element[0] + 1
		if index == input.Min {
			isValid = true
		}

		if index == input.Max {
			if isValid == true {
				isValid = false
				continue
			}
			isValid = true
		}
	}

	return isValid
}
