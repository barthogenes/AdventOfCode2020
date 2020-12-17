package day5

import "math"

// GetAnswerPart1 Get the answer for the first part of the Day 5 puzzle.
func GetAnswerPart1(input []BoardingPass) int {
	highestSeatID := 0
	for _, boardingPass := range input {
		boardingPass := calculateSeatID(boardingPass)
		highestSeatID = int(math.Max(float64(highestSeatID), float64(boardingPass.SeatID)))
	}

	return highestSeatID
}

func calculateSeatID(boardingPass BoardingPass) BoardingPass {
	boardingPass.RowNumber = calculateRowNumber(boardingPass.RowCode)
	boardingPass.ColumnNumber = calculateColumnNumber(boardingPass.ColumnCode)
	return boardingPass
}

func calculateRowNumber(rowCode string) int {
	begin, end := 0, 127
	for _, char := range rowCode {
		begin, end = binaryDivide(string(char) == "F", begin, end)
	}

	return begin
}

func calculateColumnNumber(columnCode string) int {
	begin, end := 0, 7
	for _, char := range columnCode {
		begin, end = binaryDivide(string(char) == "L", begin, end)
	}

	return begin
}

func binaryDivide(firstHalf bool, begin, end int) (newBegin, newEnd int) {
	temp := (float64(end) - float64(begin)) / 2
	if firstHalf {
		newBegin = begin
		newEnd = begin + int(temp)
	} else {
		newBegin = begin + int(math.Round(temp))
		newEnd = end
	}
	return
}
