package main

import (
    "fmt"

    "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput []string

func main() {
	inputLines := advent_helpers.ReadInput("day06/input.txt")
	
	fmt.Println("day06 Part a: ", SolveA(inputLines))
	fmt.Println("day06 Part b: ", SolveB(inputLines))
}

func ProcessInput(input []string) PuzzleInput {
    return input
}

func SolveA(input PuzzleInput) string {
	return "Not yet implemented"
}

func SolveB(input PuzzleInput) string {
	return "Not yet implemented"
}