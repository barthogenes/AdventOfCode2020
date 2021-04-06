package util

import "testing"

func Test_Open_Write_Read_And_Close_File(t *testing.T) {
	// Arrange
	file := File{}
	file.Open("file_test.txt")

	// Act
	file.Write(`abc 123
def 456`)
	got := file.Read()
	file.Close()

	// Assert
	wantWindows := "abc 123\r\ndef 456"
	wantUnix := "abc 123\ndef 456"
	if got != wantWindows && got != wantUnix {
		t.Errorf("ReadFile(input) = %v; want %v or %v", got, wantWindows, wantUnix)
	}
}
