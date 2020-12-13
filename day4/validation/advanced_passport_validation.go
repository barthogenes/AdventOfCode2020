package validation

import (
	"regexp"
	"strings"

	"github.com/barthogenes/adventofcode2020/util"
)

// AdvancedPassportValidation ...
func AdvancedPassportValidation(passport Passport) bool {
	return isBetween(passport.BirthYear, 1920, 2002) &&
		isBetween(passport.IssueYear, 2010, 2020) &&
		isBetween(passport.ExpirationYear, 2020, 2030) &&
		isValidHeight(passport.Height) &&
		isValidColor(passport.HairColor) &&
		isValidEyeColor(passport.EyeColor) &&
		isValidPassportID(passport.PassportID)
}

func isBetween(value, min, max int) bool {
	return value >= min && value <= max
}

func isValidHeight(height string) bool {
	getHeight := func(suffix string) int {
		heightWithoutSuffix := strings.TrimSuffix(height, suffix)
		if len(heightWithoutSuffix) == len(height) {
			return 0
		}
		return util.ToNumber(heightWithoutSuffix)
	}
	heightInCm := getHeight("cm")
	heightInInches := getHeight("in")
	if heightInCm > 0 && heightInCm >= 150 && heightInCm <= 193 {
		return true
	}
	if heightInInches > 0 && heightInInches >= 59 && heightInInches <= 76 {
		return true
	}

	return false
}

func isValidColor(color string) bool {
	isValid, err := regexp.MatchString("^#[\\dabcdef]{6}$", color)
	if err != nil {
		panic(err)
	}
	return isValid
}

func isValidEyeColor(eyeColor string) bool {
	switch eyeColor {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	}
	return false
}

func isValidPassportID(passportID string) bool {
	isValid, err := regexp.MatchString("^[\\d]{9}$", passportID)
	if err != nil {
		panic(err)
	}
	return isValid
}
