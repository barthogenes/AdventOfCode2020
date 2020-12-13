package validation

import (
	"testing"
)

func Test_isBetween(t *testing.T) {
	type args struct {
		value int
	}
	min := 1920
	max := 2002
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should disallow anything less than min",
			args: args{
				value: 1919,
			},
			want: false,
		},
		{
			name: "Should disallow anything greater than max",
			args: args{
				value: 2003,
			},
			want: false,
		},
		{
			name: "Should allow min",
			args: args{
				value: 1920,
			},
			want: true,
		},
		{
			name: "Should allow max",
			args: args{
				value: 2002,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBetween(tt.args.value, min, max); got != tt.want {
				t.Errorf("isBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidHeight(t *testing.T) {
	type args struct {
		height string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should not allow anything below 150cm",
			args: args{
				height: "149cm",
			},
			want: false,
		},
		{
			name: "Should not allow anything above 193cm",
			args: args{
				height: "194cm",
			},
			want: false,
		},
		{
			name: "Should not allow anything below 59in",
			args: args{
				height: "58in",
			},
			want: false,
		},
		{
			name: "Should not allow anything above 76in",
			args: args{
				height: "77in",
			},
			want: false,
		},
		{
			name: "Should end with either cm or in",
			args: args{
				height: "190",
			},
			want: false,
		},
		{
			name: "Should be equal or greater than 150cm",
			args: args{
				height: "150cm",
			},
			want: true,
		},
		{
			name: "Should be equal or smaller than 193cm",
			args: args{
				height: "193cm",
			},
			want: true,
		},
		{
			name: "Should be equal or greater than 59in",
			args: args{
				height: "59in",
			},
			want: true,
		},
		{
			name: "Should be equal or smaller than 76in",
			args: args{
				height: "76in",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidHeight(tt.args.height); got != tt.want {
				t.Errorf("isValidHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidColor(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should disallow anything that does not start with a hashtag",
			args: args{
				color: "$ffffff",
			},
			want: false,
		},
		{
			name: "Should disallow if 5 characters after the hashtag",
			args: args{
				color: "#fffff",
			},
			want: false,
		},
		{
			name: "Should disallow if 7 characters after the hashtag",
			args: args{
				color: "#fffffff",
			},
			want: false,
		},
		{
			name: "Should disallow letters after f",
			args: args{
				color: "#ghijkl",
			},
			want: false,
		},
		{
			name: "Should have exactly 6 characters after the hashtag",
			args: args{
				color: "#ffffff",
			},
			want: true,
		},
		{
			name: "Should allow numbers",
			args: args{
				color: "#123def",
			},
			want: true,
		},
		{
			name: "Should allow letters a-f",
			args: args{
				color: "#abcdef",
			},
			want: true,
		},
		{
			name: "Should allow numbers and letters",
			args: args{
				color: "#123def",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidColor(tt.args.color); got != tt.want {
				t.Errorf("isValidColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidEyeColor(t *testing.T) {
	type args struct {
		eyeColor string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should disallow pur",
			args: args{
				eyeColor: "pur",
			},
			want: false,
		},
		{
			name: "Should allow amb",
			args: args{
				eyeColor: "amb",
			},
			want: true,
		},
		{
			name: "Should allow blu",
			args: args{
				eyeColor: "blu",
			},
			want: true,
		},
		{
			name: "Should allow brn",
			args: args{
				eyeColor: "brn",
			},
			want: true,
		},
		{
			name: "Should allow gry",
			args: args{
				eyeColor: "gry",
			},
			want: true,
		},
		{
			name: "Should allow grn",
			args: args{
				eyeColor: "grn",
			},
			want: true,
		},
		{
			name: "Should allow hzl",
			args: args{
				eyeColor: "hzl",
			},
			want: true,
		},
		{
			name: "Should allow oth",
			args: args{
				eyeColor: "oth",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEyeColor(tt.args.eyeColor); got != tt.want {
				t.Errorf("isValidEyeColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPassportID(t *testing.T) {
	type args struct {
		passportID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should disallow letters",
			args: args{
				passportID: "abcdefghi",
			},
			want: false,
		},
		{
			name: "Should disallow ids with a length less than 9",
			args: args{
				passportID: "12345678",
			},
			want: false,
		},
		{
			name: "Should disallow ids with a length greater than 9",
			args: args{
				passportID: "1234567890",
			},
			want: false,
		},
		{
			name: "Should allow ids with exactly 9 digits",
			args: args{
				passportID: "123456789",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPassportID(tt.args.passportID); got != tt.want {
				t.Errorf("isValidPassportID() = %v, want %v", got, tt.want)
			}
		})
	}
}
