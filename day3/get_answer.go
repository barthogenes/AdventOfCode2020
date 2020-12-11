package day3

// GetAnswerPart1 Get the answer for the first part of the Day 3 puzzle.
func GetAnswerPart1(input [][]string) int {
	return getAnswer(input, 3, 1)
}

// GetAnswerPart2 Get the answer for the second part of the Day 3 puzzle.
func GetAnswerPart2(input [][]string) int {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	mul := 1
	for _, slope := range slopes {
		mul *= getAnswer(input, slope[0], slope[1])
	}
	return mul
}

func getAnswer(input [][]string, dx, dy int) int {
	treeCount := 0
	for x, y := 0, 0; y < len(input); y += dy {
		if string(input[y][x]) == "#" {
			treeCount++
		}
		x += dx
		if x >= len(input[0]) {
			x -= len(input[0])
		}
	}

	return treeCount
}
