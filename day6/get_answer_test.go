package day6

import "testing"

func TestGetAnswers(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "example input",
			args: args{input: []string{
				"abc",
				"a\nb\nc",
				"ab\nac",
				"a\na\na\na",
				"b",
			}},
			want:  11,
			want1: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetAnswers(tt.args.input)
			if got != tt.want {
				t.Errorf("GetAnswers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAnswers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
