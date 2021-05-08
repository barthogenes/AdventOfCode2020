package day8

import (
	"reflect"
	"testing"
)

func Test_GetAnswerForPart1(t *testing.T) {
	// Arrange
	input := `nop +0
	acc +1
	jmp +4
	acc +3
	jmp -3
	acc -99
	acc +1
	jmp -4
	acc +6`

	// Act
	got := GetAnswerForPart1(parseInput(input))

	// Assert
	want := 5
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}
