package util

import (
	"container/list"
	"testing"
)

func Test_Get_functions_support_lists_with_non_integer_items(t *testing.T) {
	// Arrange
	input := list.New()
	input.PushBack("test1")
	input.PushBack(5)
	input.PushBack(2)
	input.PushBack("test2")
	input.PushBack(8)

	// Act, Assert
	got := GetMax(input)
	want := 8
	if got != want {
		t.Errorf("GetMax() = %v, want %v", got, want)
	}
	got = GetMin(input)
	want = 2
	if got != want {
		t.Errorf("GetMin() = %v, want %v", got, want)
	}
	got = GetSum(input)
	want = 15
	if got != want {
		t.Errorf("GetSum() = %v, want %v", got, want)
	}
}
