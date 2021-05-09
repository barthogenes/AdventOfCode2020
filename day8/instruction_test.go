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

func TestInstruction_SetType_parses_NOP(t *testing.T) {
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

func TestInstruction_SetType_parses_JMP(t *testing.T) {
	// Arrange
	instruction := Instruction{}

	// Act
	instruction.SetTypeFromString("jmp")

	// Assert
	want := Jump
	if got := instruction.Type; got != want {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}

func TestInstruction_SetType_parses_ACC(t *testing.T) {
	// Arrange
	instruction := Instruction{}

	// Act
	instruction.SetTypeFromString("acc")

	// Assert
	want := Accumulate
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
