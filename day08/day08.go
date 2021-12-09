package main

import (
	"fmt"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type display struct {
	inputs  []string
	outputs []string
}
type PuzzleInput []display

func main() {
	inputLines := advent_helpers.ReadInput("day08/input.txt")

	fmt.Println("day08 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day08 Part b: ", SolveB(ProcessInput(inputLines)))
}

func ProcessInput(input []string) PuzzleInput {
	segments := make([]display, 0, len(input))
	for _, i := range input {
		split := strings.Split(i, "|")
		ins := strings.Fields(split[0])
		outs := strings.Fields(split[1])
		segments = append(segments, display{ins, outs})
	}
	return segments
}

func SolveA(input PuzzleInput) string {
	count := 0
	for _, i := range input {
		for _, j := range i.outputs {
			switch len(j) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return fmt.Sprint(count)
}

func SolveB(input PuzzleInput) string {
	count := 0
	for _, i := range input {
		mapping := getMappings(i.inputs)

		for j, output := range i.outputs {
			out := string2byte(output)
			digitPlace := advent_helpers.PowInt(10, len(i.outputs)-j-1)
			count += digitPlace * mapping[out]
		}
	}
	return fmt.Sprint(count)
}

// Convert a set of characters ('a'-'g') into a bitmask (a=1,b=2,c=4,d=8,e=16,f=32,g=64)
func string2byte(input string) (cur byte) {
	for i := 0; i < len(input); i++ {
		// 1) shift char to range 0-6 by subtracting 'a'
		// 2) convert to bitmask by shifting left that many times
		// 3) add to current accumulator
		cur = cur | (1 << (input[i] - 'a'))
	}

	return
}

// Takes a set of input strings and derives a mapping of input bitmask to expected int
func getMappings(inputs []string) (output [128]int) {
	// Maps numbers to the bitmask matching the input string
	numberBitmasks := [10]byte{}
	fives := make([]byte, 0, 3)
	sixes := make([]byte, 0, 3)

	for _, j := range inputs {
		cur := string2byte(j)
		switch len(j) {
		case 2: // 1
			numberBitmasks[1] = cur
			output[cur] = 1
		case 3: // 3
			numberBitmasks[7] = cur
			output[cur] = 7
		case 4: // 4
			numberBitmasks[4] = cur
			output[cur] = 4
		case 5: // 2, 3, 5
			fives = append(fives, cur)
		case 6: // 0, 6, 9
			sixes = append(sixes, cur)
		case 7: // 8
			numberBitmasks[8] = cur
			output[cur] = 8
		}
	}

	// Segments
	//    aa
	//   b  c
	//    dd
	//   e  f
	//    gg
	segmentsEandG := (numberBitmasks[4] | numberBitmasks[7]) ^ numberBitmasks[8] // derive the bitmask for e&g
	segmentsCandF := numberBitmasks[1]

	for _, cur := range sixes {
		hasCandF := (cur & segmentsCandF) == segmentsCandF
		hasEandG := (cur & segmentsEandG) == segmentsEandG

		if hasCandF && !hasEandG {
			numberBitmasks[9] = cur
			output[cur] = 9
		} else if hasEandG && !hasCandF {
			numberBitmasks[6] = cur
			output[cur] = 6
		} else {
			numberBitmasks[0] = cur
			output[cur] = 0
		}
	}

	segmentC := numberBitmasks[8] ^ numberBitmasks[6] // derive bitmask of segment c

	for _, cur := range fives {
		hasCandF := (cur & segmentsCandF) == segmentsCandF
		hasC := cur&segmentC > 0

		if hasCandF {
			numberBitmasks[3] = cur
			output[cur] = 3
		} else if hasC {
			numberBitmasks[2] = cur
			output[cur] = 2
		} else {
			numberBitmasks[5] = cur
			output[cur] = 5
		}
	}

	return output
}
