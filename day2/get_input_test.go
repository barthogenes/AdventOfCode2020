package day2

import "testing"

func TestStringToPasswordPolicySet(t *testing.T) {
	// Arrange
	var strInput = "1-3 a: abcde"

	// Act
	got := stringToPasswordPolicySet(strInput)

	// Assert
	want := PasswordPolicySet{
		min:      1,
		max:      3,
		char:     "a",
		password: "abcde",
	}
	if got.min != want.min {
		t.Errorf("stringToPasswordPolicySet(%s) = %d; want %d", strInput, got.min, want.min)
	}
	if got.max != want.max {
		t.Errorf("stringToPasswordPolicySet(%s) = %d; want %d", strInput, got.max, want.max)
	}
	if got.char != want.char {
		t.Errorf("stringToPasswordPolicySet(%s) = %s; want %s", strInput, got.char, want.char)
	}
	if got.password != want.password {
		t.Errorf("stringToPasswordPolicySet(%s) = %s; want %s", strInput, got.password, want.password)
	}
}
