package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput struct {
	cells []byte
	width int
}

func ProcessInput(input []string) PuzzleInput {
	width := len(input[0])
	cells := make([]byte, 0, len(input[0])*len(input))

	for _, row := range input {
		strippedRow := strings.Fields(row)[0]
		for j := range strippedRow {
			cells = append(cells, strippedRow[j]-'0')
		}
	}

	return PuzzleInput{cells, width}
}

func solve(p PuzzleInput) string {
	length := len(p.cells)
	width := p.width
	cells := p.cells

	minDistance := make([]int, length)
	visited := make(map[int]bool, length)
	toVisit := make(map[int]*ah.Record, length)
	pq := make(ah.PriorityQueue, 0, length)

	for i := 0; i < length; i++ {
		minDistance[i] = math.MaxInt
	}

	minDistance[0] = 0
	pq = append(pq, &ah.Record{Value: 0, Distance: 0})
	heap.Init(&pq)

	for len(pq) > 0 {
		next := heap.Pop(&pq).(*ah.Record)
		currentNode := next.Value

		if visited[currentNode] {
			continue
		}

		visited[currentNode] = true

		currentValue := minDistance[currentNode]

		top := currentNode - width
		if top > 0 {
			handleNode(visited, cells, minDistance, &pq, currentValue, top, toVisit)
		}

		bottom := currentNode + width
		if bottom < length {
			handleNode(visited, cells, minDistance, &pq, currentValue, bottom, toVisit)

		}

		left := currentNode - 1
		leftInRow := (currentNode % width) - 1
		if leftInRow >= 0 {
			handleNode(visited, cells, minDistance, &pq, currentValue, left, toVisit)

		}

		right := currentNode + 1
		rightInRow := (currentNode % width) + 1
		if rightInRow < width {
			handleNode(visited, cells, minDistance, &pq, currentValue, right, toVisit)
		}
	}

	return fmt.Sprint(minDistance[length-1])
}

// Copy the input 5x5 incrementing each copy and wrapping around 9
func Tile(cells []byte, origWidth int) []byte {
	newCells := make([]byte, len(cells)*25)
	newWidth := origWidth * 5
	origHeight := len(cells) / origWidth
	for j := 0; j < 5; j++ {
		for k := 0; k < 5; k++ {
			increment := byte(j + k)
			for origRow := 0; origRow < origHeight; origRow++ {
				for origCol := 0; origCol < origWidth; origCol++ {
					newRow := j*origHeight + origRow
					newCol := k*origWidth + origCol
					newCell := newRow*newWidth + newCol
					origValue := cells[origRow*origWidth+origCol]
					newValue := origValue + increment
					if newValue > 9 {
						newValue = newValue % 9
					}
					newCells[newCell] = newValue
				}
			}
		}
	}

	return newCells
}

func handleNode(visited map[int]bool, cells []byte, minDistance []int, queue *ah.PriorityQueue, baseScore int, target int, toVisit map[int]*ah.Record) {
	if !visited[target] {
		tentative := baseScore + int(cells[target])
		newVal := ah.MinInt(tentative, minDistance[target])
		minDistance[target] = newVal
		item, found := toVisit[target]
		if found {
			queue.Update(item, newVal)
		} else {
			newRecord := ah.Record{Value: target, Distance: newVal}
			heap.Push(queue, &newRecord)
			toVisit[target] = &newRecord
		}
	}
}

func SolveA(input PuzzleInput) string {
	return solve(input)
}

func SolveB(input PuzzleInput) string {
	width := input.width * 5
	cells := Tile(input.cells, input.width)

	return solve(PuzzleInput{cells, width})
}

func main() {
	inputLines := ah.ReadInput("day15/input.txt")

	fmt.Println("day15 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day15 Part b: ", SolveB(ProcessInput(inputLines)))
}
