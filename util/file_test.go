package util

import "testing"

func Test_ReadFile(t *testing.T) {
	// Act
	got := ReadFile("file_test.txt")

	// Assert
	want := "abc 123\r\ndef 456"
	if got != want {
		t.Errorf("ReadFile(input) = %v; want %v", got, want)
	}
}
