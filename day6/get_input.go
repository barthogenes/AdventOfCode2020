package day6

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input for the fourth day.
func GetInput(api api.AdventOfCodeAPI) []string {
	return parseInput(api.GetInputForDay(6))
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}
