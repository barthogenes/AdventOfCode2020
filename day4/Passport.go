package day4

import "image/color"

// Passport A struct to represent the input data from day 3.
type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	HeightInCm     int
	HairColor      color.Color
	EyeColor       color.Color
	PassportID     int
	CountryID      int
}
