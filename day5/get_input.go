package day5

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input for the fifth day.
func GetInput(api api.AdventOfCodeAPI) []BoardingPass {
	inputString := api.GetInputForDay(5)
	return parseBoardingPasses(inputString)
}

func parseBoardingPasses(input string) []BoardingPass {
	lines := strings.Split(input, "\n")
	boardingPasses := []BoardingPass{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		boardingPass := BoardingPass{}
		boardingPass.ApplySeatCode(line)
		boardingPasses = append(boardingPasses, boardingPass)
	}
	return boardingPasses
}
