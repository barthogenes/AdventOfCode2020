package day4

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input for the fourth day.
func GetInput(api api.AdventOfCodeAPI) []Passport {
	input := api.GetInputForDay(4)
	return parseInput(input)
}

func parseInput(input string) []Passport {
	passportStrings := splitIntoStrings(input)
	return parsePassports(passportStrings)
}

func splitIntoStrings(input string) []string {
	lines := strings.Split(input, "\n\n")
	for index, line := range lines {
		lines[index] = strings.ReplaceAll(line, "\n", " ")
	}

	return lines
}

func parsePassports([]string) []Passport {
	return []Passport{}
}
