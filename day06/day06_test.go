package main

import (
    "strings"
    "testing"
)

func TestA(t *testing.T) {
    testInput := ``
    expected := "N/A"

    soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

    if soln != expected {
        t.Errorf("day06 Part A doesn't match: %s != %s", soln, expected)
    }
}
func TestB(t *testing.T) {
    testInput := ``
    expected := "N/A"

    soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

    if soln != expected {
        t.Errorf("day06 Part B doesn't match: %s != %s", soln, expected)
    }
}
