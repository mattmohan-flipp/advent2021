package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	// I should use arrays, but I'm in a rush
	testInput := `start-A
    start-b
    A-c
    A-b
    b-d
    A-end
    b-end`
	testInput2 := `dc-end
    HN-start
    start-kj
    dc-start
    dc-HN
    LN-dc
    HN-end
    kj-sa
    kj-HN
    kj-dc`
	testInput3 := `fs-end
    he-DX
    fs-he
    start-DX
    pj-DX
    end-zg
    zg-sl
    zg-pj
    pj-he
    RW-he
    fs-DX
    pj-RW
    zg-RW
    start-pj
    he-WI
    zg-he
    pj-fs
    start-RW`

	expected := "10"
	expected2 := "19"
	expected3 := "226"

	soln := SolveA(ProcessInput(strings.Split(testInput, "\n")))
	soln2 := SolveA(ProcessInput(strings.Split(testInput2, "\n")))
	soln3 := SolveA(ProcessInput(strings.Split(testInput3, "\n")))

	if soln != expected {
		t.Errorf("day12 Part A Input 1 doesn't match: %s != %s", soln, expected)
	}

	if soln2 != expected2 {
		t.Errorf("day12 Part A Input 2 doesn't match: %s != %s", soln2, expected2)
	}

	if soln3 != expected3 {
		t.Errorf("day12 Part A Input 3 doesn't match: %s != %s", soln3, expected3)
	}
}
func TestB(t *testing.T) {
	testInput := `start-A
    start-b
    A-c
    A-b
    b-d
    A-end
    b-end`
	expected := "36"

	testInput3 := `dc-end
    HN-start
    start-kj
    dc-start
    dc-HN
    LN-dc
    HN-end
    kj-sa
    kj-HN
    kj-dc`
	expected3 := "103"
	soln := SolveB(ProcessInput(strings.Split(testInput, "\n")))
	soln3 := SolveB(ProcessInput(strings.Split(testInput3, "\n")))

	if soln != expected {
		t.Errorf("day12 Part B doesn't match: %s != %s", soln, expected)
	}

	if soln3 != expected3 {
		t.Errorf("day12 Part B (Input 3) doesn't match: %s != %s", soln3, expected3)
	}
}
