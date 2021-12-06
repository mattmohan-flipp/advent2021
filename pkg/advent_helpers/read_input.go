package advent_helpers

import (
	"io/ioutil"
	"log"
	"strings"
)

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
