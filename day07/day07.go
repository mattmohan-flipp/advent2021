package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput []int

func main() {
	inputLines := advent_helpers.ReadInput("day07/input.txt")

	fmt.Println("day07 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day07 Part b: ", SolveB(ProcessInput(inputLines)))
}

func calcCost1(crabs []int, target int) (cost int) {
	for _, i := range crabs {
		cost += advent_helpers.AbsInt(i - target)
	}

	return
}

func calcCost2(crabs []int, target int) (cost int) {
	for _, i := range crabs {
		diff := advent_helpers.AbsInt(i - target)
		cost += diff * (diff + 1) / 2
	}

	return
}

func ProcessInput(input []string) PuzzleInput {
	crabStrings := strings.Split(input[0], ",")
	crabs := make([]int, 0, len(crabStrings))
	for _, i := range crabStrings {
		crab, _ := strconv.Atoi(i)
		crabs = append(crabs, crab)
	}
	return crabs
}

func Solve(crabs PuzzleInput, cost func([]int, int) int) string {
	max := 0
	min := 0
	for _, i := range crabs {
		min = advent_helpers.MinInt(min, i)
		max = advent_helpers.MaxInt(max, i)
	}

	targetCost := math.MaxInt
	for i := min; i <= max; i++ {
		current := cost(crabs, i)
		if current < targetCost {
			targetCost = current
		}
	}
	return fmt.Sprint(targetCost)
}

func SolveA(crabs PuzzleInput) string {
	return Solve(crabs, calcCost1)
}

func SolveB(crabs PuzzleInput) string {
	return Solve(crabs, calcCost2)
}
