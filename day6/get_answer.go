package day6

import (
	"strings"
)

// GetAnswers Get the answer for part 1 and 2 of the Day 6 puzzle.
func GetAnswers(input []string) (int, int) {
	answerPart1 := 0
	answerPart2 := 0
	for _, group := range input {
		personCount := len(strings.Split(group, "\n"))
		allAnswers := strings.ReplaceAll(group, "\n", "")
		answerDict := make(map[rune]int)
		for _, answer := range allAnswers {
			answerDict[answer]++
		}
		answerPart1 += len(answerDict)
		for _, count := range answerDict {
			if personCount == count {
				answerPart2++
			}
		}
	}
	return answerPart1, answerPart2
}
