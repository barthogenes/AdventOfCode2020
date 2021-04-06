package day6

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	// Arrange
	var input = `abc

a
b
c

ab
ac

a
a
a
a

b`

	// Act
	got := parseInput(input)

	// Assert
	want := []string{
		"abc",
		"a\nb\nc",
		"ab\nac",
		"a\na\na\na",
		"b",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput(input) = %v; want %v", got, want)
	}
}

func Test_parseInput_handles_trailing_newline(t *testing.T) {
	// Arrange
	var input = `abc

a
b
c

ab
ac

a
a
a
a

b
`

	// Act
	got := parseInput(input)

	// Assert
	want := []string{
		"abc",
		"a\nb\nc",
		"ab\nac",
		"a\na\na\na",
		"b",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput(input) = %v; want %v", got, want)
	}
}
