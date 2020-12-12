package day4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
)

// GetInput Get the input for the fourth day.
func GetInput(api api.AdventOfCodeAPI) []Passport {
	input := api.GetInputForDay(4)
	return parseInput(input)
}

func parseInput(input string) []Passport {
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

func parsePassports(lines []string) []Passport {
	passports := []Passport{}
	for _, line := range lines {
		passport := parsePassport(line)
		passports = append(passports, passport)
	}

	return []Passport{}
}

func parsePassport(line string) Passport {
	passport := Passport{
		BirthYear: parseBirthYear(line),
	}

	return passport
}

func parseBirthYear(line string) int {
	birthYearRegexp := regexp.MustCompile("byr:(\\d+)")
	birthYearString := birthYearRegexp.FindStringSubmatch(line)[1]
	birthYear, err := strconv.Atoi(birthYearString)
	if err != nil {
		panic(err)
	}
	return birthYear
}
