package main

import (
	"fmt"
	"math"
	"strings"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput struct {
	reactions map[string]string
	start     string
}

func ProcessInput(input []string) (p PuzzleInput) {
	p.reactions = make(map[string]string, len(input)-2)
	for _, i := range input {
		row := strings.Fields(i)
		switch len(row) {
		case 1:
			p.start = row[0]
		case 3:
			p.reactions[row[0]] = row[2]
		}
	}
	return
}

func solve(input PuzzleInput, steps int) string {
	current := make(map[string]int64, len(input.reactions))
	for i := 0; i < len(input.start)-1; i++ {
		current[input.start[i:i+2]]++
	}
	for i := 0; i < steps; i++ {
		next := make(map[string]int64, len(input.reactions))
		for in, count := range current {
			inject := input.reactions[in]
			one := string(in[0]) + inject
			two := inject + string(in[1])
			next[one] += count
			next[two] += count
		}

		current = next
		total := int64(0)
		for _, j := range current {
			total += j
		}
	}

	// Only count the second half of each pair, and add in the first character
	counts := map[byte]int64{input.start[0]: 1}
	for chars, i := range current {
		counts[chars[1]] += i
	}

	// Find min/max
	max, min := int64(0), int64(math.MaxInt64)
	for _, i := range counts {
		max = ah.MaxInt64(max, i)
		min = ah.MinInt64(min, i)
	}

	return fmt.Sprint(max - min)
}

func SolveA(input PuzzleInput) string {
	return solve(input, 10)
}

func SolveB(input PuzzleInput) string {
	return solve(input, 40)
}

func main() {
	inputLines := ah.ReadInput("day14/input.txt")

	fmt.Println("day14 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day14 Part b: ", SolveB(ProcessInput(inputLines)))
}
