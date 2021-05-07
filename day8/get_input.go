package day8

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInputForPart1 Get the input for the 8th day.
func GetInputForPart1(api api.AdventOfCodeAPI) string {
	return api.GetInputForDay(8)
}

func parseInput(input string) []Instruction {
	instructions := []Instruction{}
	instructionLines := strings.Split(strings.TrimSpace(input), "\n")
	for _, instructionLine := range instructionLines {
		instruction := Instruction{}
		instruction.SetFromString(instructionLine)
		instructions = append(instructions, instruction)
	}

	return instructions
}
