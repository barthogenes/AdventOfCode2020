package day1

import (
	"strconv"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input from adventofcode.com.
func GetInput(api api.AdventOfCodeAPI) []int {
	// Get the input in bytes for the first day.
	bytes := api.GetInputInBytes(1)

	// Convert it to a string.
	inputString := string(bytes[:])

	// Split it into a Slice (dynamically sized array in Go) of strings.
	inputStringSlice := strings.Split(inputString, "\n")

	// Convert the Slice of strings to a Slice of numbers.
	inputNumberSlice := []int{}
	for index := range inputStringSlice {
		str := inputStringSlice[index]
		if str == "" {
			continue
		}

		// Convert the string to a number.
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		// Add it to the Slice of numbers.
		inputNumberSlice = append(inputNumberSlice, num)
	}

	return inputNumberSlice
}
