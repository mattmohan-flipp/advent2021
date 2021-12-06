package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2`
	expected := "5"
	input := processInput(strings.Split(testInput, "\n"))
	soln := SolveA(input)

	if soln != expected {
		t.Errorf("day05 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2`
	expected := "12"
	input := processInput(strings.Split(testInput, "\n"))

	soln := SolveB(input)

	if soln != expected {
		t.Errorf("day05 Part B doesn't match: %s != %s", soln, expected)
	}
}
