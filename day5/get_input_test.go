package day5

import (
	"reflect"
	"testing"
)

func Test_parseBoardingPasses(t *testing.T) {
	// Arrange
	var input = `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

	// Act
	got := parseBoardingPasses(input)

	// Assert
	want := []BoardingPass{
		{
			RowNumber:    44,
			ColumnNumber: 5,
			SeatID:       357,
		},
		{
			RowNumber:    70,
			ColumnNumber: 7,
			SeatID:       567,
		},
		{
			RowNumber:    14,
			ColumnNumber: 7,
			SeatID:       119,
		},
		{
			RowNumber:    102,
			ColumnNumber: 4,
			SeatID:       820,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseBoardingPasses(input) = %v; want %v", got, want)
	}
}
