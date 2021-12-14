package main

import (
	"fmt"
	"strings"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type cave struct {
	name  string
	links []string
}

func (c cave) traverse(caves PuzzleInput, path string, allowTwice bool) (count int) {
	if c.name == "end" {
		return 1
	} else if c.name == "start" && len(path) > 1 {
		return 0
	}

	newPath := path + "-" + c.name

	for _, i := range c.links {
		// If it's lower and has previously been visited then skip (unless we allow one exception)
		if ah.IsLower(i) && strings.Contains(path, i) {
			if allowTwice {
				count += caves[i].traverse(caves, newPath, false)
			}
			continue
		}

		count += caves[i].traverse(caves, newPath, allowTwice)
	}

	return
}

type PuzzleInput map[string]*cave

func (p PuzzleInput) traverse(allowTwice bool) (count int) {
	root := p["start"]
	count += root.traverse(p, "", allowTwice)
	return
}

func ProcessInput(input []string) (output PuzzleInput) {
	output = make(PuzzleInput, len(input))

	for _, i := range input {
		parts := strings.Split(i, "-")
		left := strings.Fields(parts[0])[0]
		right := strings.Fields(parts[1])[0]
		nodeLeft, foundLeft := output[left]
		nodeRight, foundRight := output[right]
		if !foundLeft {
			nodeLeft = &cave{name: left, links: make([]string, 0, len(input))}
			output[left] = nodeLeft
		}

		if !foundRight {
			nodeRight = &cave{name: right, links: make([]string, 0, len(input))}
			output[right] = nodeRight
		}
		nodeLeft.links = append(nodeLeft.links, right)

		nodeRight.links = append(nodeRight.links, left)
	}
	return output
}

func SolveA(caves PuzzleInput) string {
	return fmt.Sprint(caves.traverse(false))
}

func SolveB(caves PuzzleInput) string {
	return fmt.Sprint(caves.traverse(true))
}

func main() {
	inputLines := ah.ReadInput("day12/input.txt")

	fmt.Println("day12 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day12 Part b: ", SolveB(ProcessInput(inputLines)))
}
