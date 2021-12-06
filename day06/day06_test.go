package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `3,4,3,1,2`
	expected := "5934"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day06 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `3,4,3,1,2`
	expected := "26984457539"

	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day06 Part B doesn't match: %s != %s", soln, expected)
	}
}
