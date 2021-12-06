package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

func main() {
	inputLines := advent_helpers.ReadInput("day03/input.txt")

	bitLength := len(inputLines[0])
	inputInts := ProcessLines(inputLines)
	lineCount := len(inputInts)

	columnCounts := CountColumns(bitLength, inputInts)

	fmt.Println("Part a: ", SolveA(lineCount, columnCounts))
	fmt.Println("Part b: ", SolveB(bitLength, inputInts))
}

func SolveA(count int, columns []int) string {
	gamma, epsilon := calculateGammaEps(count, columns)
	partA := gamma * epsilon

	return strconv.FormatInt(int64(partA), 10)
}

func SolveB(bitLength int, inputInts []uint64) string {
	o2 := filterMajority(bitLength, inputInts, true)
	co2 := filterMajority(bitLength, inputInts, false)

	return strconv.FormatUint(o2[0]*co2[0], 10)
}

func filterEmpty(lines []string) (output []string) {
	for _, line := range lines {
		if len(line) > 0 {
			output = append(output, line)
		}
	}
	return
}

func ProcessLines(input []string) (output []uint64) {
	output = make([]uint64, 0, len(input))
	for _, line := range input {
		num, _ := strconv.ParseUint(line, 2, 64)
		output = append(output, num)
	}
	return
}

func CountColumns(length int, ints []uint64) (columnCounts []int) {
	columnCounts = make([]int, length)
	maxIndex := length - 1

	for _, line := range ints {
		for pos := maxIndex; pos >= 0; pos-- {
			if (line>>pos)&1 == 1 {
				columnCounts[maxIndex-pos]++
			}
		}
	}

	return
}

func filterMajority(targetBit int, ints []uint64, high bool) (selected []uint64) {
	columns := CountColumns(targetBit, ints)
	if targetBit == 0 {
		log.Fatal("Got to 0 - :(")
	}
	var criteria uint64
	if (columns[0]*2 >= len(ints) && high) || (columns[0]*2 < len(ints) && !high) {
		criteria = 1
	} else {
		criteria = 0
	}
	remaining := filter(criteria, targetBit, ints)

	if len(remaining) == 1 {
		return remaining
	}

	return filterMajority(targetBit-1, remaining, high)
}

func filter(comparator uint64, bitlength int, ints []uint64) (selected []uint64) {
	selected = make([]uint64, 0, len(ints))
	target := bitlength - 1

	for _, line := range ints {
		if (line>>target)&1 == comparator {
			selected = append(selected, line)
		}
	}
	return
}

func calculateGammaEps(count int, columns []int) (gamma uint64, epsilon uint64) {
	gamma = 0
	epsilon = 0

	for i := range columns {
		gamma = gamma << 1
		epsilon = epsilon << 1

		if columns[i]*2 > count {
			gamma++
		} else {
			epsilon++
		}
	}

	return
}
