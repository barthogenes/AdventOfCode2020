package main

import (
	"fmt"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/day1"
	day1part1 "github.com/barthogenes/adventofcode2020/day1/part1"
	day1part2 "github.com/barthogenes/adventofcode2020/day1/part2"
	"github.com/barthogenes/adventofcode2020/day2"
	day2part1 "github.com/barthogenes/adventofcode2020/day2/part1"
	day2part2 "github.com/barthogenes/adventofcode2020/day2/part2"
	"github.com/barthogenes/adventofcode2020/day3"
	"github.com/barthogenes/adventofcode2020/day4"
	"github.com/barthogenes/adventofcode2020/day5"
	"github.com/barthogenes/adventofcode2020/util"
)

func main() {
	// Getting the cookie from a text file that is not committed to GitHub.
	// If you get an error, try creating the file "cookie.txt" in the root of this repository
	// and paste the session cookie value from your "adventofcode.com" browser tab into it.
	cookie := util.ReadFile("cookie.txt")

	// Constructing the AdventOfCodeAPI class to get the input from adventofcode.com.
	api := api.AdventOfCodeAPI{Year: 2020, Cookie: cookie}

	// Showing the answers to the user.
	fmt.Println("AdventOfCode2020 puzzle implementations in Go, by Bart Hogenes :)")

	// Day 1
	fmt.Println("Day 1 - Report Repair:")
	day1Input := day1.GetInput(api)
	answerDay1Part1 := day1part1.GetAnswer(day1Input)
	fmt.Println("Part 1 answer =", answerDay1Part1)

	answerDay1Part2 := day1part2.GetAnswer(day1Input)
	fmt.Println("Part 2 answer =", answerDay1Part2)

	// Day 2
	fmt.Println("Day 2 - Password Philosophy:")
	day2Input := day2.GetInput(api)
	answerDay2Part1 := day2part1.GetAnswer(day2Input)
	fmt.Println("Part 1 answer =", answerDay2Part1)

	answerDay2Part2 := day2part2.GetAnswer(day2Input)
	fmt.Println("Part 2 answer =", answerDay2Part2)

	// Day 3
	fmt.Println("Day 3 - Toboggan Trajectory:")
	day3Input := day3.GetInput(api)
	answerDay3Part1 := day3.GetAnswerPart1(day3Input)
	fmt.Println("Part 1 answer =", answerDay3Part1)

	answerDay3Part2 := day3.GetAnswerPart2(day3Input)
	fmt.Println("Part 2 answer =", answerDay3Part2)

	// Day 4
	fmt.Println("Day 4 - Passport Processing:")
	day4Input := day4.GetInput(api)
	answerDay4Part1 := day4.GetAnswerPart1(day4Input)
	fmt.Println("Part 1 answer =", answerDay4Part1)

	answerDay4Part2 := day4.GetAnswerPart2(day4Input)
	fmt.Println("Part 2 answer =", answerDay4Part2)

	// Day 5
	fmt.Println("Day 5 - Binary Boarding:")
	day5Input := day5.GetInput(api)
	answerDay5Part1 := day5.GetAnswerPart1(day5Input)
	fmt.Println("Part 1 answer =", answerDay5Part1)

	answerDay5Part2 := day5.GetAnswerPart2(day5Input)
	fmt.Println("Part 2 answer =", answerDay5Part2)
}
