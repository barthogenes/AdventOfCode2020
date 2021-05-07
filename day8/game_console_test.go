package day8

import (
	"reflect"
	"testing"
)

func TestGameConsole_Start(t *testing.T) {
	// Arrange
	instructions := []Instruction{
		{
			Type:     NoOPeration,
			Argument: 1,
		},
	}
	gameConsole := GameConsole{}
	gameConsole.Instructions = instructions

	// Act
	gameConsole.Start()

	// Assert
	want := map[int]int{
		1: 1,
	}
	got := gameConsole.LineHitCount
	if !reflect.DeepEqual(got, want) {
		t.Errorf("gameConsole.Start() = %v, want %v", got, want)
	}
}
