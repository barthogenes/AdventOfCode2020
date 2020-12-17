package day5

import (
	"strings"

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
	lines := strings.Split(input, "\n")
	boardingPasses := []BoardingPass{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		boardingPass := parseBoardingPass(line)
		boardingPasses = append(boardingPasses, boardingPass)
	}
	return boardingPasses
}

func parseBoardingPass(line string) BoardingPass {
	return BoardingPass{
		RowCode:    line[0:7],
		ColumnCode: line[7:10],
	}
}
