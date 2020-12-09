package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/day1"
	day1part1 "github.com/barthogenes/adventofcode2020/day1/part1"
	day1part2 "github.com/barthogenes/adventofcode2020/day1/part2"
	"github.com/barthogenes/adventofcode2020/day2"
	day2part1 "github.com/barthogenes/adventofcode2020/day2/part1"
	day2part2 "github.com/barthogenes/adventofcode2020/day2/part2"
)

func main() {
	// Getting the cookie from a text file that is not committed to GitHub.
	// If you get an error, try creating the file "cookie.txt" in the root of this repository
	// and paste the session cookie value from your "adventofcode.com" browser tab into it.
	path, err := filepath.Abs("cookie.txt")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	cookie := string(data)

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
}
