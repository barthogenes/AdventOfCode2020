package part2

import (
	"testing"

	"github.com/barthogenes/adventofcode2020/day2"
)

func TestGetAnswer(t *testing.T) {
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
	if want := 1; got != want {
		t.Errorf("isValid(%v) = %v; want %v", input, got, want)
	}
}

func TestIsValid_returns_true(t *testing.T) {
	// Arrange
	var input = day2.PasswordPolicySet{
		Min:      1,
		Max:      3,
		Char:     "a",
		Password: "abcde",
	}

	// Act
	got := isValid(input)

	// Assert
	if want := true; got != want {
		t.Errorf("isValid(%v) = %v; want %v", input, got, want)
	}
}

func TestIsValid_returns_false(t *testing.T) {
	// Arrange
	var input = day2.PasswordPolicySet{
		Min:      1,
		Max:      3,
		Char:     "b",
		Password: "cdefg",
	}

	// Act
	got := isValid(input)

	// Assert
	if want := false; got != want {
		t.Errorf("isValid(%v) = %v; want %v", input, got, want)
	}
}

func TestIsValid_returns_false_again(t *testing.T) {
	// Arrange
	var input = day2.PasswordPolicySet{
		Min:      2,
		Max:      9,
		Char:     "c",
		Password: "ccccccccc",
	}

	// Act
	got := isValid(input)

	// Assert
	if want := false; got != want {
		t.Errorf("isValid(%v) = %v; want %v", input, got, want)
	}
}

func TestIsValid_returns_false_again_again(t *testing.T) {
	// Arrange
	var input = day2.PasswordPolicySet{
		Min:      1,
		Max:      3,
		Char:     "d",
		Password: "abcde",
	}

	// Act
	got := isValid(input)

	// Assert
	if want := false; got != want {
		t.Errorf("isValid(%v) = %v; want %v", input, got, want)
	}
}
