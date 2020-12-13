package day2

import (
	"regexp"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/util"
)

// GetInput Get the input from adventofcode.com.
func GetInput(api api.AdventOfCodeAPI) []PasswordPolicySet {
	// Get the input for the second day.
	inputString := api.GetInputForDay(2)

	// Convert it to an array of numbers.
	input := convertToPasswordPolicySetArray(inputString)

	return input
}

func convertToPasswordPolicySetArray(input string) []PasswordPolicySet {
	// Split into an array of strings.
	inputStrings := strings.Split(input, "\n")

	// Convert each string to a PasswordPolicySet.
	passwordPolicySets := []PasswordPolicySet{}
	for _, element := range inputStrings {
		if element == "" {
			continue
		}
		passwordPolicySet := stringToPasswordPolicySet(element)
		passwordPolicySets = append(passwordPolicySets, passwordPolicySet)
	}

	return passwordPolicySets
}

func stringToPasswordPolicySet(input string) PasswordPolicySet {
	re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
	found := re.FindStringSubmatch(input)

	min := util.ToNumber(found[1])
	max := util.ToNumber(found[2])
	char := found[3]
	password := found[4]

	return PasswordPolicySet{
		min,
		max,
		char,
		password,
	}
}
