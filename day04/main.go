package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattmohan/advent2021/pkg/advent_helpers"
)

func main() {
	inputLines := advent_helpers.ReadInput("day04/input.txt")

	fmt.Println("day4 Part a: ", SolveA(inputLines))
	fmt.Println("day4 Part b: ", SolveB(inputLines))
}

type card struct {
	nums [25]int
}

func (card card) column(j int) []int {
	column := make([]int, 5)
	for i := range column {
		column[i] = card.nums[i*5+j]
	}
	return column
}
func (card card) row(i int) []int {
	startIndex := i * 5
	return card.nums[startIndex : startIndex+5]
}

func (card card) check(called map[string]bool) bool {
	for i := 0; i < 5; i++ {
		if checkLine(called, card.column(i)) || checkLine(called, card.row(i)) {
			return true
		}
	}
	return false
}

func (card card) sumUncalled(called map[string]bool) (count int) {
	for _, i := range card.nums {
		_, found := called[strconv.Itoa(i)]
		if !found {
			count += i
		}
	}

	return
}

func (card *card) readCard(input []string) {
	for i, row := range input {
		rowStrings := strings.Fields(row)
		startIndex := i * 5
		for k, l := range rowStrings {
			idx := startIndex + k
			card.nums[idx], _ = strconv.Atoi(l)
		}
	}
}

func SolveA(input []string) string {
	numbers := strings.Split(input[0], ",")
	cards := readCards(input)

	called := make(map[string]bool, len(numbers))
	for _, num := range numbers {
		called[num] = true
		for _, card := range cards {
			if card.check(called) {
				nm, _ := strconv.Atoi(num)
				return fmt.Sprint(card.sumUncalled(called) * nm)
			}
		}
	}
	return "Not found"
}
func SolveB(input []string) string {
	numbers := strings.Split(input[0], ",")
	cards := readCards(input)

	called := make(map[string]bool, len(numbers))
	for j, num := range numbers {
		called[num] = true
		for i, card := range cards {
			if card.check(called) {
				delete(cards, i)
			}
		}
		if len(cards) == 1 {
			for _, card := range cards {
				for _, num2 := range numbers[j:] {
					called[num2] = true
					if card.check(called) {
						nm, _ := strconv.Atoi(num2)
						return fmt.Sprint(card.sumUncalled(called) * nm)
					}
				}
			}
		}
	}
	return "Not found"
}

func readCards(input []string) (cards map[int]*card) {
	cardCount := (len(input) - 1) / 6
	cards = make(map[int]*card, cardCount)

	for i := 0; i < cardCount; i++ {
		cards[i] = new(card)
		start := (i * 6) + 2
		end := start + 5
		cards[i].readCard(input[start:end])
	}

	return
}

func checkLine(called map[string]bool, row []int) bool {
	for _, i := range row {
		_, found := called[strconv.Itoa(i)]
		if !found {
			return false
		}
	}
	return true
}
