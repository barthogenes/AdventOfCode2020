package day3

import (
	"testing"
)

// Arrange
// ..##...
// #...#..
// .#....#
var input = [][]string{
	{".", ".", "#", "#", ".", ".", "."},
	{"#", ".", ".", ".", "#", ".", "."},
	{".", "#", ".", ".", ".", ".", "#"},
}

func Test_GetAnswerPart1(t *testing.T) {
	// Act
	got := GetAnswerPart1(input)

	// Assert
	if want := 1; got != want {
		t.Errorf("GetAnswer(input) = %v; want %v", got, want)
	}
}

func Test_GetAnswerPart2(t *testing.T) {
	// Act
	got := GetAnswerPart2(input)

	// Assert
	if want := 0; got != want {
		t.Errorf("GetAnswer(input) = %v; want %v", got, want)
	}
}
