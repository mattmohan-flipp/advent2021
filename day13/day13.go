package main

import (
	"fmt"
	"strconv"
	"strings"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type coord struct {
	x int
	y int
}

type fold struct {
	vert     bool
	position int
}

type PuzzleInput struct {
	dots  []coord
	folds []fold
}

func ProcessInput(input []string) (output PuzzleInput) {
	dots := make([]coord, 0, len(input))
	folds := make([]fold, 0, len(input))

	for _, i := range input {
		stripped := strings.Fields(i)
		switch len(stripped) {
		case 0:
			continue
		case 1:
			pt := strings.Split(strings.Fields(i)[0], ",")
			x, _ := strconv.Atoi(pt[0])
			y, _ := strconv.Atoi(pt[1])
			dots = append(dots, coord{x, y})
		case 3:

			fld := strings.Split(strings.Fields(i)[2], "=")
			vert := fld[0] == "y"
			position, _ := strconv.Atoi(fld[1])
			folds = append(folds, fold{vert, position})
		}
	}
	return PuzzleInput{dots, folds}
}

func SolveA(input PuzzleInput) string {
	maxX, maxY := 0, 0
	for _, i := range input.dots {
		maxX = ah.MaxInt(maxX, i.x)
		maxY = ah.MaxInt(maxY, i.y)
	}

	fold := input.folds[0]
	for i, dot := range input.dots {
		if fold.vert {
			if dot.y > fold.position {
				input.dots[i].y = 2*fold.position - dot.y
			}
		}
		if !fold.vert {
			if dot.x > fold.position {
				input.dots[i].x = 2*fold.position - dot.x
			}
		}
	}

	grid := make(map[string]bool, len(input.dots))

	for _, i := range input.dots {
		grid[fmt.Sprint(i.x, "-", i.y)] = true
	}
	return fmt.Sprint(len(grid))
}

func SolveB(input PuzzleInput) string {
	maxX, maxY := 0, 0
	for _, i := range input.dots {
		maxX = ah.MaxInt(maxX, i.x)
		maxY = ah.MaxInt(maxY, i.y)
	}

	// Set initial values
	width := maxX + 1
	height := maxY + 1

	for _, fold := range input.folds {
		for i, dot := range input.dots {
			if fold.vert {
				if dot.y > fold.position {
					height = fold.position
					input.dots[i].y = 2*fold.position - dot.y
				}
			}
			if !fold.vert {
				if dot.x > fold.position {
					width = fold.position
					input.dots[i].x = 2*fold.position - dot.x
				}
			}
		}
	}

	dotHash := make(map[string]bool, len(input.dots))

	for _, i := range input.dots {
		dotHash[fmt.Sprint(i.x, "-", i.y)] = true
	}
	out := ""
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			_, found := dotHash[fmt.Sprint(j, "-", i)]
			if found {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

func main() {
	inputLines := ah.ReadInput("day13/input.txt")

	fmt.Print("day13 Part a: \n", SolveA(ProcessInput(inputLines)), "\n")
	fmt.Print("day13 Part b: \n", SolveB(ProcessInput(inputLines)), "\n")
}
