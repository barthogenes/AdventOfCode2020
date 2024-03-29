package util

import "testing"

func Test_ToNumber_parses_string(t *testing.T) {
	// Arrange, Act
	got := ToNumber("123")

	// Assert
	if want := 123; got != want {
		t.Errorf("ToNumber(input) = %v; want %v", got, want)
	}
}

func Test_ToNumber_parses_empty_string(t *testing.T) {
	// Arrange, Act
	got := ToNumber("")

	// Assert
	if want := 0; got != want {
		t.Errorf("ToNumber(input) = %v; want %v", got, want)
	}
}

func Test_ToNumber_parses_plus_prefixed_string(t *testing.T) {
	// Arrange, Act
	got := ToNumber("+5")

	// Assert
	if want := 5; got != want {
		t.Errorf("ToNumber(input) = %v; want %v", got, want)
	}
}

func Test_ToNumber_parses_minus_prefixed_string(t *testing.T) {
	// Arrange, Act
	got := ToNumber("-5")

	// Assert
	if want := -5; got != want {
		t.Errorf("ToNumber(input) = %v; want %v", got, want)
	}
}
