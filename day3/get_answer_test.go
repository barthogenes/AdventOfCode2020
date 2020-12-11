package day3

import (
	"testing"
)

func Test_getAnswer(t *testing.T) {
	// Arrange
	// ..##...
	// #...#..
	// .#....#
	input := [][]string{
		{".", ".", "#", "#", ".", ".", "."},
		{"#", ".", ".", ".", "#", ".", "."},
		{".", "#", ".", ".", ".", ".", "#"},
	}

	// Act
	got := getAnswer(input, 3, 1)

	// Assert
	if want := 1; got != want {
		t.Errorf("GetAnswer(input) = %v; want %v", got, want)
	}
}
