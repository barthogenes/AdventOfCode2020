package day9

import "container/list"

// GetAnswerForPart1 Get the answer for part 1 of the Day 9 puzzle.
func GetAnswerForPart1(input []int, preambleSize int) int {
	xmas := XMAS{
		Preamble:     list.New(),
		PreambleSize: preambleSize,
	}

	for _, value := range input {
		result := xmas.ReadData(value)
		if result != 0 {
			return result
		}
	}

	return 0
}

// GetAnswerForPart2 Get the answer for part 2 of the Day 9 puzzle.
func GetAnswerForPart2(input []int, invalidNumber int) int {
	ewf := EncryptionWeaknessFinder{
		InvalidNumber: invalidNumber,
	}

	return ewf.FindWeakness(input)
}
