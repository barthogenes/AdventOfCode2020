package day4

import (
	"fmt"
	"reflect"
	"testing"
)

// Arrange
var input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func Test_splitIntoStrings(t *testing.T) {
	// Act
	got := splitIntoStrings(input)

	// Assert
	want := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
		"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
		"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
	}
	for _, line := range got {
		fmt.Println(line)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitIntoStrings(input) = %v; want %v", got, want)
	}
}
