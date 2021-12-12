package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `5483143223
    2745854711
    5264556173
    6141336146
    6357385478
    4167524645
    2176841721
    6882881134
    4846848554
    5283751526`
	expected := "1656"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day11 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `5483143223
    2745854711
    5264556173
    6141336146
    6357385478
    4167524645
    2176841721
    6882881134
    4846848554
    5283751526`
	expected := "195"

	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day11 Part B doesn't match: %s != %s", soln, expected)
	}
}
