package day3

import (
	"reflect"
	"testing"
)

func Test_getGrid(t *testing.T) {
	// Arrange
	var input = `
..##...
#...#..
.#....#
`

	// Act
	got := getGrid(input)

	// Assert
	want := [][]string{
		{".", ".", "#", "#", ".", ".", "."},
		{"#", ".", ".", ".", "#", ".", "."},
		{".", "#", ".", ".", ".", ".", "#"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getGrid(input) = %v; want %v", got, want)
	}
}
