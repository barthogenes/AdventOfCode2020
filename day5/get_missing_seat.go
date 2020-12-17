package day5

// GetAnswerPart2 ...
func GetAnswerPart2(input []BoardingPass) int {
	return getMissingBoardingPass(input).SeatID
}

func getMissingBoardingPass(input []BoardingPass) BoardingPass {
	missingBoardingPass := BoardingPass{}

	seats := [128][8]int{}
	for _, boardingPass := range input {
		seats[boardingPass.RowNumber][boardingPass.ColumnNumber]++
	}

	for row := 0; row < len(seats); row++ {
		// Find the row that has exactly 1 free seat.
		if !hasExactlyOneFreeSeat(seats[row]) {
			continue
		}

		for column := 0; column < len(seats[row]); column++ {
			if seats[row][column] == 0 {
				missingBoardingPass.RowNumber = row
				missingBoardingPass.ColumnNumber = column
				missingBoardingPass.SeatID = missingBoardingPass.RecalculateSeatID()
				return missingBoardingPass
			}
		}
	}

	panic("No missing boarding pass found!")
}

func hasExactlyOneFreeSeat(row [8]int) bool {
	count := 0
	for _, v := range row {
		count += v
	}
	return count == 7
}
