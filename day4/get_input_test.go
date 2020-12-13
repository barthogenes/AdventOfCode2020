package day4

import (
	"reflect"
	"testing"

	"github.com/barthogenes/adventofcode2020/day4/validation"
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
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitIntoStrings(input) = %v; want %v", got, want)
	}
}

var validPassportString = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"

func Test_parsePassport(t *testing.T) {
	// Act
	got := parsePassport(validPassportString)

	// Assert
	want := validation.Passport{
		BirthYear:      1937,
		IssueYear:      2017,
		ExpirationYear: 2020,
		Height:         "183cm",
		HairColor:      "#fffffd",
		EyeColor:       "gry",
		PassportID:     "860033327",
		CountryID:      147,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractNumber(input) = %v; want %v", got, want)
	}
}

func Test_extractValue(t *testing.T) {
	// Act
	got := extractValue(validPassportString, "byr")

	// Assert
	want := "1937"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractNumber(input) = %v; want %v", got, want)
	}
}
