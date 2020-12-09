package part1

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
	found := re.FindAllString(input.Password, -1)
	if len(found) < input.Min || len(found) > input.Max {
		return false
	}

	return true
}
