package util

import "testing"

func Test_ReadFile(t *testing.T) {
	// Act
	got := ReadFile("file_test.txt")

	// Assert
	wantWindows := "abc 123\r\ndef 456"
	wantUnix := "abc 123\ndef 456"
	if got != wantWindows && got != wantUnix {
		t.Errorf("ReadFile(input) = %v; want %v or %v", got, wantWindows, wantUnix)
	}
}
