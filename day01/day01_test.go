package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `199
    200
    208
    210
    200
    207
    240
    269
    260
    263`
	expected := "7"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day01 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `199
    200
    208
    210
    200
    207
    240
    269
    260
    263`
	expected := "5"

	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day01 Part B doesn't match: %s != %s", soln, expected)
	}
}
