package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput []string

func ProcessInput(input []string) PuzzleInput {
	out := make([]string, 0, len(input))
	for _, i := range input {
		out = append(out, strings.Fields(i)[0])
	}
	return out
}

const (
	Round int = iota
	Square
	Curly
	Angled
)

func parseOut(in string) string {
	pairs := map[byte]byte{'{': '}', '[': ']', '<': '>', '(': ')'}

	for i := 1; i < len(in); i++ {
		if pairs[in[i-1]] == in[i] {
			return parseOut(in[:i-1] + in[i+1:])
		}
	}
	return in
}

func scoreStr(in string) (score int) {
	scores := map[rune]int{'}': 1197, ']': 57, ')': 3, '>': 25137}

	for _, ch := range in {
		pnts, found := scores[ch]
		if found {
			return pnts
		}
	}

	return 0
}

func SolveA(input PuzzleInput) string {
	score := 0
	for _, line := range input {
		parsed := parseOut(line)

		score += scoreStr(parsed)
	}
	return fmt.Sprint(score)
}

func SolveB(input PuzzleInput) string {
	scoreMap := map[byte]int{'}': 3, ']': 2, ')': 1, '>': 4}
	pairs := map[byte]byte{'{': '}', '[': ']', '<': '>', '(': ')'}
	scores := make([]int, 0, len(input))

	for _, line := range input {
		parsed := parseOut(line)
		if scoreStr(parsed) == 0 {
			lineScore := 0
			for i := len(parsed) - 1; i >= 0; i-- {
				// Reverse and flip brackets
				lineScore = lineScore*5 + scoreMap[pairs[parsed[i]]]
			}
			scores = append(scores, lineScore)
		}
	}
	sort.Ints(scores)
	return fmt.Sprint(scores[len(scores)/2])
}

func main() {
	inputLines := advent_helpers.ReadInput("day10/input.txt")

	fmt.Println("day10 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day10 Part b: ", SolveB(ProcessInput(inputLines)))
}
