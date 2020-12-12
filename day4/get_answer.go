package day4

// GetAnswer Get the answer for the Day 4 puzzle.
func GetAnswer(input []Passport) int {
	validPassportsCount := 0
	for _, passport := range input {
		if isValid(passport) {
			validPassportsCount++
		}
	}

	return validPassportsCount
}

func isValid(passport Passport) bool {
	return passport.BirthYear != 0 &&
		passport.IssueYear != 0 &&
		passport.ExpirationYear != 0 &&
		passport.Height != "" &&
		passport.HairColor != "" &&
		passport.EyeColor != "" &&
		passport.PassportID != ""
}
