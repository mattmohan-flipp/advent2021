package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `16,1,2,0,4,2,7,1,2,14`
	expected := "37"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day07 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `16,1,2,0,4,2,7,1,2,14`
	expected := "168"

	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day07 Part B doesn't match: %s != %s", soln, expected)
	}
}
