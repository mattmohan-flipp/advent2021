package main

import (
	"testing"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

func TestA(t *testing.T) {
	file := advent_helpers.ReadInput("../input.txt")
	inputInts := ProcessLines(file)

	columns := CountColumns(len(file[0]), inputInts)
	expected := "198"
	sln := SolveA(len(inputInts), columns)
	if sln != expected {
		t.Errorf("Part A doesn't match: %s != %s", sln, expected)
	}
}
func TestB(t *testing.T) {
	file := advent_helpers.ReadInput("../input.txt")
	inputInts := ProcessLines(file)

	sln := SolveB(len(file[0]), inputInts)
	expected := "230"

	if sln != expected {
		t.Errorf("Part B doesn't match: %s != %s", sln, expected)
	}
}
