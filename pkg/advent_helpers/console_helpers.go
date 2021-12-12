package advent_helpers

import "fmt"

const ConsConsoleSafeChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/-=`~*@&^%$#!(){}[];:<>,.?\\"

func Byte2ConsoleChar(in byte) string {
	return string(ConsConsoleSafeChars[in%byte(len(ConsConsoleSafeChars))])
}
func RenderByteGrid(input [][]byte, hideZero bool) (out string) {
	for _, row := range input {
		for _, cell := range row {
			if cell == 0 && hideZero {
				out += " "
			} else {
				char := Byte2ConsoleChar(cell)
				colour := cell/byte(len(ConsConsoleSafeChars)) + 1
				out += fmt.Sprintf("\033[3%dm%v\033[0m", colour, string(char))
			}
		}
		out += "\n "
	}

	return
}
