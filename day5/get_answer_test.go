package day5

import "testing"

func TestGetAnswerPart1(t *testing.T) {
	type args struct {
		input []BoardingPass
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Return the highest SeatID",
			args: args{
				input: []BoardingPass{
					{
						SeatID: 123,
					},
					{
						SeatID: 456,
					},
					{
						SeatID: 789,
					},
				},
			},
			want: 789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAnswerPart1(tt.args.input); got != tt.want {
				t.Errorf("GetAnswerPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
