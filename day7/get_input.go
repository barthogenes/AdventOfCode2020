package day7

import (
	"regexp"
	"strings"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/util"
)

// GetInputForPart1 Get the input for the 7th day.
func GetInputForPart1(api api.AdventOfCodeAPI) InvertedBagList {
	return parseInputForPart1(api.GetInputForDay(7))
}

func parseInputForPart1(input string) InvertedBagList {
	bagList := initializeInvertedBagList(input)
	for bag, fitsXTimesIn := range bagList {
		bagFitsXTimesInRegex := regexp.MustCompile(`(\w* \w*) bags contain (\d \w* \w* bags?, )*(\d) ` + bag)
		matches := bagFitsXTimesInRegex.FindAllStringSubmatch(input, -1)
		for _, match := range matches {
			fitsXTimesIn[match[1]] = util.ToNumber(match[3])
		}
	}
	return bagList
}

func initializeInvertedBagList(input string) InvertedBagList {
	bagList := InvertedBagList{}
	bagLines := strings.Split(strings.TrimSpace(input), "\n")
	bagRegex := regexp.MustCompile(`(\w* \w*) bags contain`)
	for _, bagLine := range bagLines {
		bagList[bagRegex.FindStringSubmatch(bagLine)[1]] = make(FitsXTimesIn)
	}

	return bagList
}

// GetInputForPart2 Get the input for the 7th day.
func GetInputForPart2(api api.AdventOfCodeAPI) BagList {
	return parseInputForPart2(api.GetInputForDay(7))
}

func parseInputForPart2(input string) BagList {
	bagList := BagList{}
	bagLines := strings.Split(strings.TrimSpace(input), "\n")
	bagRegex := regexp.MustCompile(`^(\w* \w*)`)
	containsRegex := regexp.MustCompile(`(\d) (\w* \w*)`)
	for _, bagLine := range bagLines {
		bag := bagRegex.FindStringSubmatch(bagLine)[1]
		bagList[bag] = make(Contains)
		containedBagLines := containsRegex.FindAllStringSubmatch(bagLine, -1)
		for _, containedBagLine := range containedBagLines {
			bagList[bag][containedBagLine[2]] = util.ToNumber(containedBagLine[1])
		}
	}

	return bagList
}
