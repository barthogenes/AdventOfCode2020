package day5

import "math"

// BoardingPass ...
type BoardingPass struct {
	RowNumber    int
	ColumnNumber int
	SeatID       int
}

// ApplySeatCode ...
func (boardingPass *BoardingPass) ApplySeatCode(seatCode string) {
	boardingPass.RowNumber = calculateRowNumber(seatCode[0:7])
	boardingPass.ColumnNumber = calculateColumnNumber(seatCode[7:10])
	boardingPass.SeatID = boardingPass.RecalculateSeatID()
}

// RecalculateSeatID ...
func (boardingPass *BoardingPass) RecalculateSeatID() int {
	return boardingPass.RowNumber*8 + boardingPass.ColumnNumber
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
