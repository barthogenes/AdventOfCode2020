package day1

import (
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/util"
)

// GetInput Get input from adventofcode.com.
func GetInput(api api.AdventOfCodeAPI) []int {
	// Get the input for the first day.
	inputString := api.GetInputForDay(1)

	// Convert it to an array of numbers.
	input := convertToNumbersArray(inputString)

	return input
}

func convertToNumbersArray(inputString string) []int {
	// Split into a array of strings.
	inputStrings := strings.Split(inputString, "\n")

	inputNumbers := []int{}
	for _, element := range inputStrings {
		if element == "" {
			continue
		}

		// Parse the string to a number.
		num := util.ToNumber(element)

		// Add it to the array of numbers.
		inputNumbers = append(inputNumbers, num)
	}

	return inputNumbers
}
