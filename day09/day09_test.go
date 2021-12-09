package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `2199943210
    3987894921
    9856789892
    8767896789
    9899965678`
	expected := "15"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day09 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `2199943210
    3987894921
    9856789892
    8767896789
    9899965678`
	expected := "1134"

	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day09 Part B doesn't match: %s != %s", soln, expected)
	}
}
