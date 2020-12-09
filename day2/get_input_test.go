package day2

import (
	"reflect"
	"testing"
)

func TestStringToPasswordPolicySet(t *testing.T) {
	// Arrange
	var input = `
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

	// Act
	got := convertToPasswordPolicySetArray(input)

	// Assert
	want := []PasswordPolicySet{
		{
			Min:      1,
			Max:      3,
			Char:     "a",
			Password: "abcde",
		},
		{
			Min:      1,
			Max:      3,
			Char:     "b",
			Password: "cdefg",
		},
		{
			Min:      2,
			Max:      9,
			Char:     "c",
			Password: "ccccccccc",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("convertToPasswordPolicySetArray(%v) = %v; want %v", input, got, want)
	}
}
