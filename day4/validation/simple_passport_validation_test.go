package validation

import (
	"testing"
)

func Test_SimplePassportValidation(t *testing.T) {
	type args struct {
		input Passport
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Passport 1",
			args: args{
				input: Passport{
					BirthYear:      1937,
					IssueYear:      2017,
					ExpirationYear: 2020,
					Height:         "183cm",
					HairColor:      "#fffffd",
					EyeColor:       "gry",
					PassportID:     "860033327",
					CountryID:      147,
				},
			},
			want: true,
		},
		{
			name: "Passport 2",
			args: args{
				input: Passport{
					BirthYear:      1929,
					IssueYear:      2013,
					ExpirationYear: 2023,
					Height:         "",
					HairColor:      "#cfa07d",
					EyeColor:       "amb",
					PassportID:     "028048884",
					CountryID:      350,
				},
			},
			want: false,
		},
		{
			name: "Passport 3",
			args: args{
				input: Passport{
					BirthYear:      1931,
					IssueYear:      2013,
					ExpirationYear: 2024,
					Height:         "179cm",
					HairColor:      "#ae17e1",
					EyeColor:       "brn",
					PassportID:     "760753108",
					CountryID:      0,
				},
			},
			want: true,
		},
		{
			name: "Passport 4",
			args: args{
				input: Passport{
					BirthYear:      0,
					IssueYear:      2011,
					ExpirationYear: 2025,
					Height:         "59in",
					HairColor:      "#cfa07d",
					EyeColor:       "brn",
					PassportID:     "166559648",
					CountryID:      0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplePassportValidation(tt.args.input); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
