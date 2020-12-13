package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/day4/validation"
	"github.com/barthogenes/adventofcode2020/util"
)

// GetInput Get the input for the fourth day.
func GetInput(api api.AdventOfCodeAPI) []validation.Passport {
	input := api.GetInputForDay(4)
	return parseInput(input)
}

func parseInput(input string) []validation.Passport {
	passportStrings := splitIntoStrings(input)
	return parsePassports(passportStrings)
}

func splitIntoStrings(input string) []string {
	lines := strings.Split(input, "\n\n")
	for index, line := range lines {
		lines[index] = strings.ReplaceAll(line, "\n", " ")
	}

	return lines
}

func parsePassports(lines []string) []validation.Passport {
	passports := []validation.Passport{}
	for _, line := range lines {
		passport := parsePassport(line)
		passports = append(passports, passport)
	}

	return passports
}

func parsePassport(line string) validation.Passport {
	passport := validation.Passport{
		BirthYear:      util.ToNumber(extractValue(line, "byr")),
		IssueYear:      util.ToNumber(extractValue(line, "iyr")),
		ExpirationYear: util.ToNumber(extractValue(line, "eyr")),
		Height:         extractValue(line, "hgt"),
		HairColor:      extractValue(line, "hcl"),
		EyeColor:       extractValue(line, "ecl"),
		PassportID:     extractValue(line, "pid"),
		CountryID:      util.ToNumber(extractValue(line, "cid")),
	}

	return passport
}

func extractValue(line, prefix string) string {
	regex := regexp.MustCompile(fmt.Sprintf("%s:(#?[\\w|\\d]+)", prefix))
	matches := regex.FindStringSubmatch(line)
	if len(matches) == 2 {
		return matches[1]
	}

	return ""
}
