package day8

import (
	"reflect"
	"testing"
)

func TestInstruction_SetFromString(t *testing.T) {
	// Arrange
	instruction := Instruction{}

	// Act
	instruction.SetFromString("nop +15")

	// Assert
	want := Instruction{
		Type:     NoOPeration,
		Argument: 15,
	}
	if !reflect.DeepEqual(instruction, want) {
		t.Errorf("parseInput() = %v, want %v", instruction, want)
	}
}

func TestInstruction_SetType(t *testing.T) {
	// Arrange
	instruction := Instruction{}

	// Act
	instruction.SetTypeFromString("nop")

	// Assert
	want := NoOPeration
	if got := instruction.Type; got != want {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}

func TestInstruction_SetArgument(t *testing.T) {
	// Arrange
	instruction := Instruction{}

	// Act
	instruction.SetArgumentFromString("-5")

	// Assert
	want := NoOPeration
	if got := instruction.Type; got != want {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}
