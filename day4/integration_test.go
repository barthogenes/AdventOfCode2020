package day4

import (
	"reflect"
	"testing"
)

func Test_GetAnswers(t *testing.T) {
	// Arrange
	var inputString = `
iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	// Act
	input := parseInput(inputString)
	got := GetAnswer(input)

	// Assert
	want := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAnswer(input) = %v; want %v", got, want)
	}
}