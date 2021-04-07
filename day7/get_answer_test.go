package day7

import "testing"

func TestGetAnswerForPart1(t *testing.T) {
	type args struct {
		input InvertedBagList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example input",
			args: args{input: InvertedBagList{
				"light red":   FitsXTimesIn{},
				"dark orange": FitsXTimesIn{},
				"bright white": FitsXTimesIn{
					"light red":   1,
					"dark orange": 3,
				},
				"muted yellow": FitsXTimesIn{
					"light red":   2,
					"dark orange": 4,
				},
				"shiny gold": FitsXTimesIn{
					"bright white": 1,
					"muted yellow": 2,
				},
				"dark olive": FitsXTimesIn{
					"shiny gold": 1,
				},
				"vibrant plum": FitsXTimesIn{
					"shiny gold": 2,
				},
				"faded blue": FitsXTimesIn{
					"muted yellow": 9,
					"dark olive":   3,
					"vibrant plum": 5,
				},
				"dotted black": FitsXTimesIn{
					"dark olive":   4,
					"vibrant plum": 6,
				},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAnswerForPart1(tt.args.input)
			if got != tt.want {
				t.Errorf("GetAnswerForPart1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAnswerForPart2(t *testing.T) {
	type args struct {
		input BagList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example input 1",
			args: args{input: BagList{
				"light red": Contains{
					"bright white": 1,
					"muted yellow": 2,
				},
				"dark orange": Contains{
					"bright white": 3,
					"muted yellow": 4,
				},
				"bright white": Contains{
					"shiny gold": 1,
				},
				"muted yellow": Contains{
					"shiny gold": 2,
				},
				"shiny gold": Contains{
					"dark olive":   1,
					"vibrant plum": 2,
				},
				"dark olive": Contains{
					"faded blue":   3,
					"dotted black": 4,
				},
				"vibrant plum": Contains{
					"faded blue":   5,
					"dotted black": 6,
				},
				"faded blue":   Contains{},
				"dotted black": Contains{},
			}},
			want: 32,
		},
		{
			name: "example input 2",
			args: args{input: BagList{
				"shiny gold": Contains{
					"dark red": 2,
				},
				"dark red": Contains{
					"dark orange": 2,
				},
				"dark orange": Contains{
					"dark yellow": 2,
				},
				"dark yellow": Contains{
					"dark green": 2,
				},
				"dark green": Contains{
					"dark blue": 2,
				},
				"dark blue": Contains{
					"dark violet": 2,
				},
				"dark violet": Contains{},
			}},
			want: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAnswerForPart2(tt.args.input)
			if got != tt.want {
				t.Errorf("GetAnswerForPart2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
