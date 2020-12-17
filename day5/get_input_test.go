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
			RowCode:      "FBFBBFF",
			RowNumber:    0,
			ColumnCode:   "RLR",
			ColumnNumber: 0,
			SeatID:       0,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseBoardingPasses(input) = %v; want %v", got, want)
	}
}
