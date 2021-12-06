package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cbroglie/mustache"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Define args
	day := flag.Int("day", 0, "Which day to generate")

	// Parse ARGV
	flag.Parse()

	// Generate a new folder
	dayString := fmt.Sprintf("day%02d", *day)

	err := os.Mkdir(dayString, 0755)
	check(err)

	mainContents, err := mustache.RenderFile("generate/main.mustache", map[string]string{"dayString": dayString})
	check(err)

	testContents, err := mustache.RenderFile("generate/test.mustache", map[string]string{"dayString": dayString})
	check(err)

	os.WriteFile(dayString+"/"+dayString+".go", []byte(mainContents), 0644)
	os.WriteFile(dayString+"/"+dayString+"_test.go", []byte(testContents), 0644)
	os.WriteFile(dayString+"/input.txt", []byte(""), 0644)
}
