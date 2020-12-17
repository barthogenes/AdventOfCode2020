package day5

import (
	"testing"
)

// Arrange
var input = []BoardingPass{
	{
		RowCode:      "FBFBBFF",
		RowNumber:    0,
		ColumnCode:   "RLR",
		ColumnNumber: 0,
		SeatID:       0,
	},
	{
		RowCode:      "BFFFBBF",
		RowNumber:    0,
		ColumnCode:   "RRR",
		ColumnNumber: 0,
		SeatID:       0,
	},
	{
		RowCode:      "FFFBBBF",
		RowNumber:    0,
		ColumnCode:   "RRR",
		ColumnNumber: 0,
		SeatID:       0,
	},
	{
		RowCode:      "BBFFBBF",
		RowNumber:    0,
		ColumnCode:   "RLL",
		ColumnNumber: 0,
		SeatID:       0,
	},
}

func Test_GetAnswerPart1(t *testing.T) {
	// Act
	got := GetAnswerPart1(input)

	// Assert
	if want := 820; got != want {
		t.Errorf("GetAnswerPart1(input) = %v; want %v", got, want)
	}
}
