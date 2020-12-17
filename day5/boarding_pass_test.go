package day5

import (
	"testing"
)

func Test_calculateRowNumber(t *testing.T) {
	type args struct {
		rowCode string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "RowCode FBFBBFF",
			args: args{
				rowCode: "FBFBBFF",
			},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRowNumber(tt.args.rowCode); got != tt.want {
				t.Errorf("calculateRowNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateColumnNumber(t *testing.T) {
	type args struct {
		columnCode string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ColumnCode RLR",
			args: args{
				columnCode: "RLR",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateColumnNumber(tt.args.columnCode); got != tt.want {
				t.Errorf("calculateColumnNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryDivide(t *testing.T) {
	type args struct {
		firstHalf bool
		begin     int
		end       int
	}
	tests := []struct {
		name         string
		args         args
		wantNewBegin int
		wantNewEnd   int
	}{
		{
			name: "Boarding pass FBFBBFFRLR letter 1 (F)",
			args: args{
				firstHalf: true,
				begin:     0,
				end:       127,
			},
			wantNewBegin: 0,
			wantNewEnd:   63,
		},
		{
			name: "Boarding pass FBFBBFFRLR letter 2 (B)",
			args: args{
				firstHalf: false,
				begin:     0,
				end:       63,
			},
			wantNewBegin: 32,
			wantNewEnd:   63,
		},
		{
			name: "Boarding pass FBFBBFFRLR letter 3 (F)",
			args: args{
				firstHalf: true,
				begin:     32,
				end:       63,
			},
			wantNewBegin: 32,
			wantNewEnd:   47,
		},
		{
			name: "Boarding pass FBFBBFFRLR letter 4 (B)",
			args: args{
				firstHalf: false,
				begin:     32,
				end:       47,
			},
			wantNewBegin: 40,
			wantNewEnd:   47,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewBegin, gotNewEnd := binaryDivide(tt.args.firstHalf, tt.args.begin, tt.args.end)
			if gotNewBegin != tt.wantNewBegin {
				t.Errorf("binaryDivide() gotNewBegin = %v, want %v", gotNewBegin, tt.wantNewBegin)
			}
			if gotNewEnd != tt.wantNewEnd {
				t.Errorf("binaryDivide() gotNewEnd = %v, want %v", gotNewEnd, tt.wantNewEnd)
			}
		})
	}
}

func TestBoardingPass_RecalculateSeatID(t *testing.T) {
	tests := []struct {
		name         string
		boardingPass *BoardingPass
		want         int
	}{
		{
			name: "Boarding pass FBFBBFFRLR",
			boardingPass: &BoardingPass{
				RowNumber:    44,
				ColumnNumber: 5,
			},
			want: 357,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.boardingPass.RecalculateSeatID(); got != tt.want {
				t.Errorf("BoardingPass.RecalculateSeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
