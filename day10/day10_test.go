package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	testInput := `[({(<(())[]>[[{[]{<()<>>
        [(()[<>])]({[<{<<[]>>(
        {([(<{}[<>[]}>{[]{[(<()>
        (((({<>}<{<{<>}{[]{[]{}
        [[<[([]))<([[{}[[()]]]
        [{[{({}]{}}([{[{{{}}([]
        {<[[]]>}<{[{[{[]{()[[[]
        [<(<(<(<{}))><([]([]()
        <{([([[(<>()){}]>(<<{{
        <{([{{}}[<[[[<>{}]]]>[]]`
	expected := "26397"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day10 Part A doesn't match: %s != %s", soln, expected)
	}
}
func TestB(t *testing.T) {
	testInput := `[({(<(())[]>[[{[]{<()<>>
        [(()[<>])]({[<{<<[]>>(
        {([(<{}[<>[]}>{[]{[(<()>
        (((({<>}<{<{<>}{[]{[]{}
        [[<[([]))<([[{}[[()]]]
        [{[{({}]{}}([{[{{{}}([]
        {<[[]]>}<{[{[{[]{()[[[]
        [<(<(<(<{}))><([]([]()
        <{([([[(<>()){}]>(<<{{
        <{([{{}}[<[[[<>{}]]]>[]]`
	expected := "288957"

	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))

	if soln != expected {
		t.Errorf("day10 Part B doesn't match: %s != %s", soln, expected)
	}
}
