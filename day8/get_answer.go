package day8

// GetAnswerForPart1 Get the answer for part 1 of the Day 8 puzzle.
func GetAnswerForPart1(input []Instruction) int {
	gameConsole := GameConsole{
		Instructions: input,
	}

	acc, _ := gameConsole.Start()
	return acc
}

// GetAnswerForPart2 Get the answer for part 2 of the Day 8 puzzle.
func GetAnswerForPart2(input []Instruction) int {
	for index := range input {
		inputCopy := make([]Instruction, len(input))
		copy(inputCopy, input)
		if inputCopy[index].Type == NoOPeration {
			inputCopy[index].Type = Jump
		} else if inputCopy[index].Type == Jump {
			inputCopy[index].Type = NoOPeration
		}

		gameConsole := GameConsole{
			Instructions: inputCopy,
		}

		acc, returnCode := gameConsole.Start()

		if returnCode == 0 {
			return acc
		}
	}

	return -1
}
