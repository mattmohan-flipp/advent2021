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
			t.Errorf("day18 Part A (input %d) doesn't match: %s != %s", i, soln, testcase.Out)
		}
	}

	reduce := []string{
		"[[[[4,3],4],4],[7,[[8,4],9]]]\n[1,1]",
		"[1,1]\n[2,2]\n[3,3]\n[4,4]",
		"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]",
		"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]\n[6,6]",
	}
	reduceExpected := []string{
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
	}

	for i := range reduce {
		input := strings.Split(reduce[i], "\n")
		soln := ParseSnailNumber(input[0])
		for j := 1; j < len(input); j++ {
			next := ParseSnailNumber(input[j])
			soln = soln.addNode(next)
			soln.Reduce()
		}

		if soln.String() != reduceExpected[i] {
			t.Errorf("day18 Part A (input %d) doesn't match:\nsolution\n----\n%s\n\nexpected\n----\n%s\n", i, soln.String(), reduceExpected[i])
		}
	}
}

func TestB(t *testing.T) {
	testcases := ah.ReadTests("b")
	for i, testcase := range testcases {
		soln := SolveB(ProcessInput(testcase.In))

		if soln != strings.Join(testcase.Out, "\n") {
			t.Errorf("day18 Part B (input %d) doesn't match: %s != %s", i, soln, testcase.Out)
		}
	}
}
