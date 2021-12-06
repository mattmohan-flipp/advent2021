package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

func main() {
	splitInput := advent_helpers.ReadInput("day02/input.txt")

	depth := 0
	distance := 0

	depth2 := 0
	distance2 := 0
	aim2 := 0

	for _, line := range splitInput {
		parts := strings.Split(line, " ")
		dir := parts[0]
		len, _ := strconv.Atoi(parts[1])
		switch dir {
		case "forward":
			distance += len

			distance2 += len
			depth2 += aim2 * len
		case "up":
			depth -= len

			aim2 -= len
		case "down":
			depth += len

			aim2 += len
		}
	}

	fmt.Println(distance*depth, distance2*depth2)
}
