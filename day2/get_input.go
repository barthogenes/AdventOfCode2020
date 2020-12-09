package day2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// PasswordPolicySet A struct to represent the input data from day 2.
type PasswordPolicySet struct {
	min      int
	max      int
	char     string
	password string
}

// GetInput Get the input from adventofcode.com.
func GetInput(api api.AdventOfCodeAPI) []PasswordPolicySet {
	// Get the input in bytes for the second day.
	bytes := api.GetInputInBytes(2)

	// Convert it to a string.
	inputString := string(bytes[:])

	// Split it into a Slice (dynamically sized array in Go) of strings.
	inputStringSlice := strings.Split(inputString, "\n")

	// Convert the Slice of strings to a Slice of PasswordPolicySets.
	passwordPolicySets := []PasswordPolicySet{}
	for index := range inputStringSlice {
		str := inputStringSlice[index]
		if str == "" {
			continue
		}

		passwordPolicySet := stringToPasswordPolicySet(str)

		// Add it to the Slice.
		passwordPolicySets = append(passwordPolicySets, passwordPolicySet)
	}

	return passwordPolicySets
}

func stringToPasswordPolicySet(input string) PasswordPolicySet {
	re := regexp.MustCompile("(\\d)-(\\d) (\\w): (\\w+)")
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
