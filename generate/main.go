package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func check(e error) {
	if e != nil {
		log.Error().Err(e).Msg("")
		os.Exit(1)
	}
}

func main() {
	// Setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Get the current day of the month
	curDay := time.Now().Day()
	curYear := time.Now().Year()

	sessionEnv, _ := os.LookupEnv("AOC_SESSION")

	// Define args
	day := flag.Int("day", curDay, "Which day to generate")
	year := flag.Int("year", curYear, "Which year to generate")
	sessionKey := flag.String("session", sessionEnv, "Your session cookie")
	quiet := flag.Bool("quiet", false, "sets log level to info")

	// Parse ARGV
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if *quiet {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Debug().Int("day", *day).Int("year", *year).Bool("debug", *quiet).Msg("Running script")

	dayString := fmt.Sprintf("day%02d", *day)

	inputChan := make(chan io.ReadCloser)

	if len(*sessionKey) > 10 {
		log.Debug().Msg("Session key found. Making http request...")
		go MakeAocRequest(sessionKey, *year, *day, inputChan)
	} else {
		log.Debug().Msg("No session key found. Mocking http request...")
		go MockAocRequest(sessionKey, 0, 0, inputChan)
	}

	log.Debug().Str("folder", dayString).Msg("Creating folder")
	err := os.Mkdir(dayString, 0755)
	check(err)

	mainContents, err := mustache.RenderFile("generate/main.mustache", map[string]string{"dayString": dayString})
	check(err)

	testContents, err := mustache.RenderFile("generate/test.mustache", map[string]string{"dayString": dayString})
	check(err)

	mainFilename := dayString + "/" + dayString + ".go"
	log.Debug().Str("contents", mainContents).Str("filename", mainFilename).Msg("Writing file " + mainFilename)
	os.WriteFile(mainFilename, []byte(mainContents), 0644)

	testFilename := dayString + "/" + dayString + "_test.go"
	log.Debug().Str("contents", testContents).Str("filename", testFilename).Msg("Writing file " + testFilename)
	os.WriteFile(testFilename, []byte(testContents), 0644)

	log.Debug().Msg("Opening file")
	file, err := os.OpenFile(dayString+"/input.txt", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	check(err)
	defer file.Close()

	inputResponse := <-inputChan
	defer inputResponse.Close()

	io.Copy(file, inputResponse)
}

func MockAocRequest(sessionKey *string, year int, day int, reqChan chan io.ReadCloser) {
	reqChan <- io.NopCloser(strings.NewReader(""))
}

func MakeAocRequest(sessionKey *string, year int, day int, reqChan chan io.ReadCloser) {
	client := &http.Client{Timeout: time.Second * 10}

	// Generate a request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	check(err)

	// Handle Auth
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = *sessionKey
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	check(err)

	if resp.StatusCode != 200 {
		errMsg, _ := ioutil.ReadAll(resp.Body)
		log.Error().Bytes("message", errMsg).Msg("Error downloading input from AoC")
		os.Exit(2)
	}

	reqChan <- resp.Body
}
