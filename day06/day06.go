package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput []int

func main() {
	inputLines := advent_helpers.ReadInput("day06/input.txt")

	fmt.Println("day06 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day06 Part b: ", SolveB(ProcessInput(inputLines)))
}

func ProcessInput(input []string) PuzzleInput {
	splitInput := strings.Split(input[0], ",")
	ints := make([]int, 0, len(splitInput))
	for _, i := range splitInput {
		number, _ := strconv.Atoi(i)
		ints = append(ints, number)
	}

	return ints
}

func SolveA(fish PuzzleInput) string {
	return Solve(fish, 80)
}

func SolveB(fish PuzzleInput) string {
	return Solve(fish, 256)
}

func Solve(fish PuzzleInput, days int) string {
	sum := 0
	tally := [9]int{}

	for _, i := range fish {
		tally[i]++
	}

	for i := 0; i < days; i++ {
		newFish := tally[0]
		for j := 1; j <= 8; j++ {
			tally[j-1] = tally[j]
		}
		tally[8] = 0
		tally[6] += newFish
		tally[8] += newFish
	}
	for _, i := range tally {
		sum += i
	}

	return fmt.Sprint(sum)
}
