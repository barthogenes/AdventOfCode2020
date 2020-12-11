package day4

import (
	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input for the fourth day.
func GetInput(api api.AdventOfCodeAPI) []Passport {
	return parseInput(api.GetInputForDay(4))
}

func parseInput(input string) []Passport {
	return []Passport{}
}
