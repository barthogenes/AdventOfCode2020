package day5

import "math"

// GetAnswerPart1 Get the answer for the first part of the Day 5 puzzle.
func GetAnswerPart1(input []BoardingPass) int {
	highestSeatID := 0
	for _, boardingPass := range input {
		highestSeatID = int(math.Max(float64(highestSeatID), float64(boardingPass.SeatID)))
	}

	return highestSeatID
}
