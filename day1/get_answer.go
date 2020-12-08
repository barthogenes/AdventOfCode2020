package day1

// GetAnswerPart1 Get the answer for first part of the Day 1 Puzzle.
func GetAnswerPart1(input []int) int {
	for x := range input {
		for y := range input {
			result := input[x] + input[y]

			if result == 2020 {
				return input[x] * input[y]
			}
		}
	}

	panic("Error: No solution was found!")
}

// GetAnswerPart2 Get the answer for the second part of the Day 1 Puzzle.
func GetAnswerPart2(input []int) int {
	for x := range input {
		for y := range input {
			for z := range input {
				result := input[x] + input[y] + input[z]

				if result == 2020 {
					return input[x] * input[y] * input[z]
				}
			}
		}
	}

	panic("Error: No solution was found!")
}
