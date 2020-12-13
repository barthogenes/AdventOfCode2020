package validation

// SimplePassportValidation ...
func SimplePassportValidation(passport Passport) bool {
	return passport.BirthYear != 0 &&
		passport.IssueYear != 0 &&
		passport.ExpirationYear != 0 &&
		passport.Height != "" &&
		passport.HairColor != "" &&
		passport.EyeColor != "" &&
		passport.PassportID != ""
}
