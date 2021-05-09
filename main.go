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
	"github.com/barthogenes/adventofcode2020/day6"
	"github.com/barthogenes/adventofcode2020/day7"
	"github.com/barthogenes/adventofcode2020/day8"
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
	defer util.Track(util.Runningtime("execution"))

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

	// Day 6
	fmt.Println("Day 6 - Custom Customs:")
	day6Input := day6.GetInput(api)
	answerDay6Part1, answerDay6Part2 := day6.GetAnswers(day6Input)
	fmt.Println("Part 1 answer =", answerDay6Part1)
	fmt.Println("Part 2 answer =", answerDay6Part2)

	// Day 7
	fmt.Println("Day 7 - Handy Haversacks:")
	day7Part1Input := day7.GetInputForPart1(api)
	day7Part2Input := day7.GetInputForPart2(api)
	answerDay7Part1 := day7.GetAnswerForPart1(day7Part1Input)
	answerDay7Part2 := day7.GetAnswerForPart2(day7Part2Input)
	fmt.Println("Part 1 answer =", answerDay7Part1)
	fmt.Println("Part 2 answer =", answerDay7Part2)

	// Day 8
	fmt.Println("Day 8 - Handheld Halting:")
	day8Input := day8.GetInput(api)
	answerDay8Part1 := day8.GetAnswerForPart1(day8Input)
	answerDay8Part2 := day8.GetAnswerForPart2(day8Input)
	fmt.Println("Part 1 answer =", answerDay8Part1)
	fmt.Println("Part 2 answer =", answerDay8Part2)
}
