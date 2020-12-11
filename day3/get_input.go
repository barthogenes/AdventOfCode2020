package day3

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input for the third day.
func GetInput(api api.AdventOfCodeAPI) [][]string {
	inputString := api.GetInputForDay(3)
	return getGrid(inputString)
}

func getGrid(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := [][]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}
