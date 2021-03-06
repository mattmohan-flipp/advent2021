package main

import (
    "strings"
    "testing"

    ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

func TestA(t *testing.T) {
	testcases := ah.ReadTests("a")
	for i, testcase := range testcases {
		soln := SolveA(ProcessInput(testcase.In))

		if soln != strings.Join(testcase.Out, "\n") {
			t.Errorf("day13 Part A (input %d) doesn't match: %s != %s", i, soln, testcase.Out)
		}
	}
}

func TestB(t *testing.T) {
	testcases := ah.ReadTests("b")
	for i, testcase := range testcases {
		soln := SolveB(ProcessInput(testcase.In))

		if soln != strings.Join(testcase.Out, "\n") {
			t.Errorf("day13 Part B (input %d) doesn't match: %s != %s", i, soln, testcase.Out)
		}
	}
}
