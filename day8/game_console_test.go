package day8

import (
	"reflect"
	"testing"
)

func TestGameConsole_Start_handles_jumps_correctly(t *testing.T) {
	// Arrange
	instructions := []Instruction{
		{
			Type:     Jump,
			Argument: 3,
		},
		{
			Type:     NoOPeration,
			Argument: 0,
		},
		{
			Type:     Accumulate,
			Argument: 2,
		},
		{
			Type:     Jump,
			Argument: -2,
		},
	}
	gameConsole := GameConsole{}
	gameConsole.Instructions = instructions

	// Act
	gameConsole.Start()

	// Assert
	want := map[int]int{
		1: 1,
		4: 1,
		2: 1,
		3: 1,
	}
	got := gameConsole.LineHitCount
	if !reflect.DeepEqual(got, want) {
		t.Errorf("gameConsole.LineHitCount = %v, want %v", got, want)
	}
}

func TestGameConsole_Start_returns_Accumulator(t *testing.T) {
	// Arrange
	instructions := []Instruction{
		{
			Type:     Accumulate,
			Argument: 25,
		},
	}
	gameConsole := GameConsole{}
	gameConsole.Instructions = instructions

	// Act
	got := gameConsole.Start()

	// Assert
	if want := 25; got != want {
		t.Errorf("gameConsole.Start() = %v, want %v", got, want)
	}
}

func TestGameConsole_Start_returns_Accumulator_on_infinite_loop_detection(t *testing.T) {
	// Arrange
	instructions := []Instruction{
		{
			Type:     Accumulate,
			Argument: 1,
		},
		{
			Type:     Jump,
			Argument: -1,
		},
	}
	gameConsole := GameConsole{}
	gameConsole.Instructions = instructions

	// Act
	got := gameConsole.Start()

	// Assert
	if want := 1; got != want {
		t.Errorf("gameConsole.Start() = %v, want %v", got, want)
	}
}
