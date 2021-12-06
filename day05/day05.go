package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type point struct {
	x int
	y int
}
type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (line line) horizVert() bool {
	return line.horiz() || line.vert()
}

func (line line) horiz() bool {
	return line.y1 == line.y2
}

func (line line) vert() bool {
	return line.x1 == line.x2
}

func (line line) steps() (xStep int, yStep int) {
	if line.x1 > line.x2 {
		xStep = -1
	} else if line.x1 < line.x2 {
		xStep = 1
	}
	if line.y1 > line.y2 {
		yStep = -1
	} else if line.y1 < line.y2 {
		yStep = 1
	}

	return
}

func (line line) length() (length int) {
	lengthX := maxInt(line.x1, line.x2) - minInt(line.x1, line.x2)
	lengthY := maxInt(line.y1, line.y2) - minInt(line.y1, line.y2)
	length = maxInt(lengthX, lengthY)

	return
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxInt(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

type PuzzleInput []line

func main() {
	inputLines := advent_helpers.ReadInput("day05/input.txt")
	lines := processInput(inputLines)

	fmt.Println("day05 Part a: ", SolveA(lines))
	fmt.Println("day05 Part b: ", SolveB(lines))
}

func processInput(inputLines []string) (lines PuzzleInput) {
	lines = make([]line, 0, len(inputLines))
	for _, i := range inputLines {
		row := strings.Fields(i)
		start := strings.Split(row[0], ",")
		end := strings.Split(row[2], ",")

		x1, _ := strconv.Atoi(start[0])
		x2, _ := strconv.Atoi(end[0])
		y1, _ := strconv.Atoi(start[1])
		y2, _ := strconv.Atoi(end[1])

		lines = append(lines, line{x1: x1, x2: x2, y1: y1, y2: y2})
	}

	return
}

func SolveA(input PuzzleInput) string {
	lines := make([]line, 0, len(input))

	// filter to only horizontal/vertical
	for _, line := range input {
		if line.horizVert() {
			lines = append(lines, line)
		}
	}

	blocks := make(map[point]uint8)

	for _, line := range lines {
		xStep, yStep := line.steps()
		length := line.length()
		x := line.x1
		y := line.y1

		for i := 0; i <= length; i++ {
			blocks[point{x, y}]++
			x += xStep
			y += yStep
		}
	}

	overlapCount := 0
	for _, i := range blocks {
		if i >= 2 {
			overlapCount++
		}
	}
	return fmt.Sprint(overlapCount)

}
func SolveB(input PuzzleInput) string {

	blocks := make(map[point]uint8)

	for _, line := range input {
		xStep, yStep := line.steps()
		length := line.length()
		x := line.x1
		y := line.y1

		for i := 0; i <= length; i++ {
			blocks[point{x, y}]++
			x += xStep
			y += yStep
		}
	}

	overlapCount := 0
	for _, i := range blocks {
		if i >= 2 {
			overlapCount++
		}
	}

	return fmt.Sprint(overlapCount)
}
