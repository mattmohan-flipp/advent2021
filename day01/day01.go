package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput []int

func ProcessInput(input []string) PuzzleInput {
	nums := make([]int, 0, len(input))
	for _, i := range input {
		num, _ := strconv.Atoi(strings.Fields(i)[0])
		nums = append(nums, num)
	}
	return nums
}

func SolveA(input PuzzleInput) string {
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}
	return fmt.Sprint(count)
}

func SolveB(input PuzzleInput) string {
	count := 0
	for i := 3; i < len(input); i++ {
		if input[i]+input[i-1]+input[i-2] > input[i-1]+input[i-2]+input[i-3] {
			count++
		}
	}
	return fmt.Sprint(count)
}

func main() {
	inputLines := advent_helpers.ReadInput("day01/input.txt")

	fmt.Println("day01 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day01 Part b: ", SolveB(ProcessInput(inputLines)))
}
