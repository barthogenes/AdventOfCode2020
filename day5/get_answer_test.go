package day5

import (
	"reflect"
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

func Test_calculateSeatID(t *testing.T) {
	type args struct {
		boardingPass BoardingPass
	}
	tests := []struct {
		name string
		args args
		want BoardingPass
	}{
		{
			name: "Boarding pass FBFBBFFRLR",
			args: args{
				boardingPass: BoardingPass{
					RowCode:      "FBFBBFF",
					RowNumber:    0,
					ColumnCode:   "RLR",
					ColumnNumber: 0,
					SeatID:       0,
				},
			},
			want: BoardingPass{
				RowCode:      "FBFBBFF",
				RowNumber:    44,
				ColumnCode:   "RLR",
				ColumnNumber: 5,
				SeatID:       357,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSeatID(tt.args.boardingPass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
