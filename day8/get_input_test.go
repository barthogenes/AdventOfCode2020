package day8

import (
	"reflect"
	"testing"
)

func Test_parseInput_parses_NOP_JMP_ACC(t *testing.T) {
	// Arrange
	input := `nop +0
jmp +1
acc +2`

	// Act
	got := parseInput(input)

	// Assert
	want := []Instruction{
		{
			Type:     NoOPeration,
			Argument: 0,
		},
		{
			Type:     Jump,
			Argument: 1,
		},
		{
			Type:     Accumulate,
			Argument: 2,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}
