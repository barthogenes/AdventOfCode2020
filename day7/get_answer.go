package day7

// GetAnswerForPart1 Get the answer for part 1 of the Day 7 puzzle.
func GetAnswerForPart1(input InvertedBagList) int {
	bagsThatCanHoldBag := make(map[string]int)
	fillMapOfBagsThatCanHoldBag("shiny gold", input, bagsThatCanHoldBag)
	return len(bagsThatCanHoldBag)
}

func fillMapOfBagsThatCanHoldBag(bag string, input InvertedBagList, bagsThatCanHoldBag map[string]int) {
	fitsXTimesIn := input[bag]
	for b := range fitsXTimesIn {
		bagsThatCanHoldBag[b] = 1
		fillMapOfBagsThatCanHoldBag(b, input, bagsThatCanHoldBag)
	}
}

// GetAnswerForPart2 Get the answer for part 2 of the Day 7 puzzle.
func GetAnswerForPart2(input BagList) int {
	bagCount := 0
	backCanHold("shiny gold", input, 1, &bagCount)
	return bagCount
}

func backCanHold(bag string, input BagList, multiplier int, bagCount *int) {
	contains := input[bag]
	for b, count := range contains {
		*bagCount += count * multiplier
		backCanHold(b, input, multiplier*count, bagCount)
	}
}
