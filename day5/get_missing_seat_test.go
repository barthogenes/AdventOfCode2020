package day5

import (
	"reflect"
	"testing"
)

func Test_getMissingBoardingPass(t *testing.T) {
	type args struct {
		input []BoardingPass
	}
	tests := []struct {
		name string
		args args
		want BoardingPass
	}{
		{
			name: "Find missing board pass",
			args: args{
				input: []BoardingPass{
					{
						ColumnNumber: 0,
					},
					{
						ColumnNumber: 1,
					},
					{
						ColumnNumber: 2,
					},
					{
						ColumnNumber: 3,
					},
					{
						ColumnNumber: 4,
					},
					{
						ColumnNumber: 6,
					},
					{
						ColumnNumber: 7,
					},
				},
			},
			want: BoardingPass{
				ColumnNumber: 5,
				SeatID:       5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMissingBoardingPass(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMissingBoardingPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
