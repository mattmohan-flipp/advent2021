package advent_helpers

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TestInput struct {
	In  []string
	Out []string
}

func ReadInput(fname string) []string {
	input, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}
	split := strings.Split(string(input), "\n")
	if split[len(split)-1] == "" {
		return split[:len(split)-1]
	}
	return split
}

func ReadTests(part string) (testcases []TestInput) {
	files, err := os.ReadDir("testcases")
	CheckError(err)

	testcases = make([]TestInput, 0, len(files))
	for _, i := range files {
		if !i.Type().IsRegular() {
			continue
		}
		splt := strings.Split(i.Name(), "_")
		if splt[1] == "in" && strings.Contains(splt[2], part) {
			input, err := os.ReadFile("testcases/" + i.Name())
			CheckError(err)

			output, err := os.ReadFile("testcases/" + splt[0] + "_out_" + part + ".txt")
			CheckError(err)

			testcases = append(testcases, TestInput{
				strings.Split(string(input), "\n"),
				strings.Split(string(output), "\n"),
			})
		}
	}

	return
}
