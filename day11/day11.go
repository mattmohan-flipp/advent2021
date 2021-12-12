package main

import (
	"fmt"
	"strings"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

const (
	colLength int = 10
	rowLength int = 10
)

type PuzzleInput []byte

func (p *PuzzleInput) runStepA() (changed int) {
	// Part1 - Increment each cell
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			cell := i*rowLength + j

			(*p)[cell]++
		}
	}

	// Part2 - Flash if >9 and then repeat until we complete a loop without finding a new flasher
	flashes := make([]bool, colLength*rowLength)
	for found := true; found; {
		found = false
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				cell := i*rowLength + j

				if (*p)[cell] > 9 && !flashes[cell] {
					// Mark this as flashed and ensure that the loop will run at least once more
					flashes[cell] = true
					found = true

					// Define range to increment
					startI, endI := ah.MaxInt(i-1, 0), ah.MinInt(i+1, 9)
					startJ, endJ := ah.MaxInt(j-1, 0), ah.MinInt(j+1, 9)

					// Increment neighbors
					for x := startJ; x <= endJ; x++ {
						for y := startI; y <= endI; y++ {
							(*p)[y*rowLength+x]++
						}
					}
				}
			}
		}
	}

	// Part 3 - If flashed drop to 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			cell := i*rowLength + j
			if (*p)[cell] > 9 {
				(*p)[cell] = 0
				changed++
			}
		}
	}

	return
}

func ProcessInput(input []string) PuzzleInput {
	out := make([]byte, 0, len(input)*rowLength)
	for _, i := range input {
		row := strings.Fields(i)[0]
		for j := range row {
			out = append(out, row[j]-byte('0'))
		}
	}
	return out
}

func SolveA(grid PuzzleInput) string {
	total := 0

	for step := 0; step < 100; step++ {
		flashes := grid.runStepA()
		total += flashes
	}

	return fmt.Sprint(total)
}

func SolveB(grid PuzzleInput) string {
	for step := 0; true; step++ {
		grid.runStepA()
		found := false
		for i := 0; i < colLength*rowLength; i++ {
			if grid[i] > 0 {
				found = true
			}
		}
		if !found {
			return fmt.Sprint(step + 1)
		}

	}

	return "Error - This follows an infinite loop. How did we get here?!?"
}

func main() {
	inputLines := ah.ReadInput("day11/input.txt")

	fmt.Println("day11 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day11 Part b: ", SolveB(ProcessInput(inputLines)))
}
