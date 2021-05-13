package day9

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/util"
)

// GetInput Get the input for the 8th day.
func GetInput(api api.AdventOfCodeAPI) []int {
	return parseInput(api.GetInputForDay(9))
}

func parseInput(input string) []int {
	data := []int{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		data = append(data, util.ToNumber(strings.TrimSpace(line)))
	}

	return data
}
