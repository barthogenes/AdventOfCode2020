package day9

import "testing"

func Test_GetAnswerForPart1(t *testing.T) {
	// Arrange
	input := `35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576`

	// Act
	got := GetAnswerForPart1(parseInput(input), 5)

	// Assert
	want := 127
	if got != want {
		t.Errorf("GetAnswerForPart1() = %v, want %v", got, want)
	}
}
