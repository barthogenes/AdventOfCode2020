package day7

import (
	"reflect"
	"testing"
)

func Test_parseInputForPart1(t *testing.T) {
	// Arrange
	var input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

	// Act
	got := parseInputForPart1(input)

	// Assert
	want := InvertedBagList{
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
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput(input) = %v; want %v", got, want)
	}
}

func Test_parseInputForPart2_example1(t *testing.T) {
	// Arrange
	var input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

	// Act
	got := parseInputForPart2(input)

	// Assert
	want := BagList{
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
			"faded blue": 9,
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
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput(input) = %#v; want %#v", got, want)
	}
}

func Test_parseInputForPart2_example2(t *testing.T) {
	// Arrange
	var input = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`

	// Act
	got := parseInputForPart2(input)

	// Assert
	want := BagList{
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
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput(input) = %#v; want %#v", got, want)
	}
}
