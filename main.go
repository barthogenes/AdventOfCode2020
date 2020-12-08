package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/barthogenes/adventofcode2020/api"
	"github.com/barthogenes/adventofcode2020/day1"
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
	fmt.Println("Day 1:")
	day1Input := day1.GetInput(api)
	answerDay1Part1 := day1.GetAnswerPart1(day1Input)
	fmt.Println("Part 1 answer =", answerDay1Part1)

	answerDay1Part2 := day1.GetAnswerPart2(day1Input)
	fmt.Println("Part 2 answer =", answerDay1Part2)
}
