package day4

import "github.com/barthogenes/adventofcode2020/day4/validation"

// GetAnswerPart1 Get the answer for part 1 of the Day 4 puzzle.
func GetAnswerPart1(input []validation.Passport) int {
	return validatePassports(input, validation.SimplePassportValidation)
}

// GetAnswerPart2 Get the answer for part 2 of the Day 4 puzzle.
func GetAnswerPart2(input []validation.Passport) int {
	return validatePassports(input, validation.AdvancedPassportValidation)
}

func validatePassports(input []validation.Passport, validatorFunc func(passport validation.Passport) bool) int {
	validPassportsCount := 0
	for _, passport := range input {
		if validatorFunc(passport) {
			validPassportsCount++
		}
	}

	return validPassportsCount
}
