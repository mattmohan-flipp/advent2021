package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput [][]byte

func ProcessInput(input []string) PuzzleInput {
	ints := make([][]byte, 0, len(input))
	for i, row := range input {
		cleanRow := strings.Fields(row)[0]
		ints = append(ints, make([]byte, len(cleanRow)))
		for j := range cleanRow {
			ints[i][j] = cleanRow[j] - '0'
		}
	}
	return ints
}

func SolveA(input PuzzleInput) string {
	count := 0
	for i := range input {
		for j := range input[i] {
			top := i-1 < 0 || input[i-1][j] > input[i][j]
			bottom := i+1 == len(input) || input[i+1][j] > input[i][j]
			left := j-1 < 0 || input[i][j-1] > input[i][j]
			right := j+1 == len(input[i]) || input[i][j+1] > input[i][j]
			if top && bottom && left && right {
				count += int(1 + input[i][j])
			}
		}
	}
	return fmt.Sprint(count)
}

func SolveB(input PuzzleInput) string {
	answer := 1
	colour := byte(1)

	checked := make([][]byte, len(input))
	topRegions := make([]int, 3)

	for i := range input {
		checked[i] = make([]byte, len(input[i]))
	}

	for i := range input {
		for j := range input[i] {
			if checked[i][j] > 0 || input[i][j] == 9 {
				continue
			}
			result := checkAdjacent(input, i, j, checked, colour)
			if result > topRegions[0] {
				topRegions[0] = result
				sort.Ints(topRegions)
			}
			colour++
		}
	}

	fmt.Println("\n", printGrid(checked))

	for _, i := range topRegions {
		answer *= i
	}

	return fmt.Sprint(answer)
}

func checkAdjacent(grid [][]byte, i int, j int, checked [][]byte, colour byte) (count int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || checked[i][j] > 0 || grid[i][j] == 9 {
		return 0
	}

	checked[i][j] = colour

	top := checkAdjacent(grid, i-1, j, checked, colour)
	bottom := checkAdjacent(grid, i+1, j, checked, colour)
	left := checkAdjacent(grid, i, j-1, checked, colour)
	right := checkAdjacent(grid, i, j+1, checked, colour)

	return top + bottom + left + right + 1
}

func main() {
	inputLines := advent_helpers.ReadInput("day09/input.txt")

	fmt.Println("day09 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day09 Part b: ", SolveB(ProcessInput(inputLines)))
}

func printGrid(input [][]byte) (out string) {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/-*"
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 0 {
				out += " "
			} else {
				char := chars[input[i][j]%66]
				colour := int(input[i][j])/len(chars) + 1
				out += fmt.Sprintf("\033[3%dm%v\033[0m", colour, string(char))
			}
		}
		out += "\n "
	}

	return
}
