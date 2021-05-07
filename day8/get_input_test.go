package day8

import (
	"reflect"
	"testing"
)

func Test_parseInput_parses_NOP(t *testing.T) {
	// Arrange
	input := "nop +0"

	// Act
	got := parseInput(input)

	// Assert
	want := []Instruction{
		{
			Type:     NoOPeration,
			Argument: 0,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}
