package day8

import (
	"reflect"
	"testing"
)

func Test_GetAnswerForPart1(t *testing.T) {
	// Arrange
	input := []Instruction{
		{
			Type:     Accumulate,
			Argument: 3,
		},
		{
			Type:     Jump,
			Argument: -1,
		},
	}

	// Act
	got := GetAnswerForPart1(input)

	// Assert
	want := 3
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}
