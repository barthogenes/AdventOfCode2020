package day1

import "testing"

// Note: A general scaffold for tests can be autogenerated by VSCode using right click -> "Go: Generate Unit Tests For Function".
// I am using my own test structure here.

// Arrange
var input = []int{1721, 979, 366, 299, 675, 1456}

func TestGetAnswerPart1(t *testing.T) {
	// Act
	got := GetAnswerPart1(input)

	// Assert
	want := 514579
	if got != want {
		t.Errorf("GetAnswerPart1(input) = %d; want %d", got, want)
	}
}

func TestGetAnswerPart2(t *testing.T) {
	// Act
	got := GetAnswerPart2(input)

	// Assert
	want := 241861950
	if got != want {
		t.Errorf("GetAnswerPart1(input) = %d; want %d", got, want)
	}
}
