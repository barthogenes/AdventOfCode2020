package day3

import (
	"reflect"
	"testing"
)

var inputString = `..##.........##.........##.........##.........##.........##.......
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#`

func Test_GetAnswers(t *testing.T) {
	type args struct {
		dx, dy int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Arrange
		{
			name: "Slope [1, 1]",
			args: args{
				dx: 1,
				dy: 1,
			},
			want: 2,
		},
		{
			name: "Slope [3, 1]",
			args: args{
				dx: 3,
				dy: 1,
			},
			want: 7,
		},
		{
			name: "Slope [5, 1]",
			args: args{
				dx: 5,
				dy: 1,
			},
			want: 3,
		},
		{
			name: "Slope [7, 1]",
			args: args{
				dx: 7,
				dy: 1,
			},
			want: 4,
		},
		{
			name: "Slope [1, 2]",
			args: args{
				dx: 1,
				dy: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act, Assert
			input := getGrid(inputString)
			if got := getAnswer(input, tt.args.dx, tt.args.dy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAnswer(input) = %v, want %v", got, tt.want)
			}
		})
	}
}
