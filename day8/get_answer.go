package day8

// GetAnswerForPart1 Get the answer for part 1 of the Day 8 puzzle.
func GetAnswerForPart1(input []Instruction) int {
	gameConsole := GameConsole{
		Instructions: input,
	}

	return gameConsole.Start()
}
