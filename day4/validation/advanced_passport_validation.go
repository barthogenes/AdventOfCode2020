package validation

// AdvancedPassportValidation ...
func AdvancedPassportValidation(passport Passport) bool {
	return isBetween(passport.BirthYear, 1920, 2002) &&
		isBetween(passport.IssueYear, 2010, 2020) &&
		isBetween(passport.ExpirationYear, 2020, 2030) &&
		isValidHeight(passport.Height) &&
		isValidColor(passport.HairColor) &&
		eyeColorExists(passport.EyeColor) &&
		isValidPassportID(passport.PassportID)
}

func isBetween(value, min, max int) bool {
	return false
}

func isValidHeight(height string) bool {
	return false
}

func isValidColor(color string) bool {
	return false
}

func eyeColorExists(eyeColor string) bool {
	return false
}

func isValidPassportID(passportID string) bool {
	return false
}
