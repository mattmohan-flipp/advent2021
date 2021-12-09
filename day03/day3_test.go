package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	file := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	inputs := strings.Split(file, "\n")
	inputInts := ProcessLines(inputs)

	columns := CountColumns(len(inputs[0]), inputInts)
	expected := "198"
	sln := SolveA(len(inputInts), columns)
	if sln != expected {
		t.Errorf("Part A doesn't match: %s != %s", sln, expected)
	}
}
func TestB(t *testing.T) {
	file := `00100
	11110
	10110
	10111
	10101
	01111
	00111
	11100
	10000
	11001
	00010
	01010`
	inputs := strings.Split(file, "\n")
	inputInts := ProcessLines(inputs)

	sln := SolveB(len(inputs[0]), inputInts)
	expected := "230"

	if sln != expected {
		t.Errorf("Part B doesn't match: %s != %s", sln, expected)
	}
}
