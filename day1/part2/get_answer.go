package part2

// GetAnswer Get the answer for the second part of the Day 1 Puzzle.
func GetAnswer(input []int) int {
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
