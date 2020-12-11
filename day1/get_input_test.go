package day1

import (
	"reflect"
	"testing"
)

func Test_convertToNumbersArray(t *testing.T) {
	// Arrange
	var input = `
1721
979
366
299
675
1456
`

	// Act
	got := convertToNumbersArray(input)

	// Assert
	want := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("convertToNumbers(%v) = %v; want %v", input, got, want)
	}
}
