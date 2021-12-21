package main

import (
	"strings"
	"testing"
	"time"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

func TestA(t *testing.T) {
	testcases := ah.ReadTests("a")
	for i, testcase := range testcases {
		timeout := time.After(100 * time.Millisecond)
		solnCh := make(chan string)
		go func() {
			solnCh <- SolveA(ProcessInput(testcase.In))
		}()
		select {
		case <-timeout:
			t.Errorf("Timeout on part %d", i)
		case soln := <-solnCh:
			if soln != strings.Join(testcase.Out, "\n") {
				t.Errorf("day16 Part A (input %d) doesn't match: %s != %s", i, soln, testcase.Out)
			}
		}
	}
}

func TestB(t *testing.T) {
	testcases := ah.ReadTests("b")
	for i, testcase := range testcases {
		soln := SolveB(ProcessInput(testcase.In))

		if soln != strings.Join(testcase.Out, "\n") {
			t.Errorf("day16 Part B (input %d) doesn't match: %s != %s", i, soln, testcase.Out)
		}
	}
}
