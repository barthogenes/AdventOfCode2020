package day5

import (
	"github.com/barthogenes/adventofcode2020/api"
)

// BoardingPass ...
type BoardingPass struct {
	RowCode      string
	RowNumber    int
	ColumnCode   string
	ColumnNumber int
	SeatID       int
}

// GetInput Get the input for the fifth day.
func GetInput(api api.AdventOfCodeAPI) []BoardingPass {
	inputString := api.GetInputForDay(5)
	return parseBoardingPasses(inputString)
}

func parseBoardingPasses(input string) []BoardingPass {
	return []BoardingPass{}
}
