package part1

// GetAnswer Get the answer for first part of the Day 1 Puzzle.
func GetAnswer(input []int) int {
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
