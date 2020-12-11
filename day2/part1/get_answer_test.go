package part1

import (
	"testing"

	"github.com/barthogenes/adventofcode2020/day2"
)

func Test_GetAnswer(t *testing.T) {
	// Arrange
	var input = []day2.PasswordPolicySet{
		{
			Min:      1,
			Max:      3,
			Char:     "a",
			Password: "abcde",
		},
		{
			Min:      1,
			Max:      3,
			Char:     "b",
			Password: "cdefg",
		},
		{
			Min:      2,
			Max:      9,
			Char:     "c",
			Password: "ccccccccc",
		},
	}

	// Act
	got := GetAnswer(input)

	// Assert
	if want := 2; got != want {
		t.Errorf("isValid(%v) = %v; want %v", input, got, want)
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		input day2.PasswordPolicySet
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isValid_1",
			args: args{
				input: day2.PasswordPolicySet{
					Min:      1,
					Max:      3,
					Char:     "a",
					Password: "abcde",
				},
			},
			want: true,
		},
		{
			name: "Test_isValid_2",
			args: args{
				input: day2.PasswordPolicySet{
					Min:      1,
					Max:      3,
					Char:     "b",
					Password: "cdefg",
				},
			},
			want: false,
		},
		{
			name: "Test_isValid_3",
			args: args{
				input: day2.PasswordPolicySet{
					Min:      2,
					Max:      9,
					Char:     "c",
					Password: "ccccccccc",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.input); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
