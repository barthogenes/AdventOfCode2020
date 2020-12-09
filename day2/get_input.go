package day2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// PasswordPolicySet A struct to represent the input data from day 2.
type PasswordPolicySet struct {
	Min      int
	Max      int
	Char     string
	Password string
}

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

	min, err := strconv.Atoi(found[1])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(found[2])
	if err != nil {
		panic(err)
	}

	char := found[3]
	password := found[4]

	return PasswordPolicySet{
		min,
		max,
		char,
		password,
	}
}
