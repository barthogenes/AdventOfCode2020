package util

import (
	"container/list"
	"testing"
)

func TestGetMax(t *testing.T) {
	// Arrange
	input := list.New()
	input.PushFront(5)
	input.PushFront(2)
	input.PushFront(8)

	// Act
	got := GetMax(input)

	// Assert
	want := 8
	if got != want {
		t.Errorf("GetMax() = %v, want %v", got, want)
	}
}

func TestGetMin(t *testing.T) {
	// Arrange
	input := list.New()
	input.PushFront(5)
	input.PushFront(2)
	input.PushFront(8)

	// Act
	got := GetMin(input)

	// Assert
	want := 2
	if got != want {
		t.Errorf("GetMax() = %v, want %v", got, want)
	}
}

func TestGetSum(t *testing.T) {
	// Arrange
	input := list.New()
	input.PushFront(5)
	input.PushFront(2)
	input.PushFront(8)

	// Act
	got := GetSum(input)

	// Assert
	want := 15
	if got != want {
		t.Errorf("GetMax() = %v, want %v", got, want)
	}
}
