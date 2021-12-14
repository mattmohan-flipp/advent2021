package advent_helpers

import "fmt"

const ConsConsoleSafeChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/-=`~*@&^%$#!(){}[];:<>,.?\\"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func Byte2ConsoleChar(in byte) string {
	return string(ConsConsoleSafeChars[in%byte(len(ConsConsoleSafeChars))])
}
func RenderBytesLinear(input []byte, zeroChar string, width int, colourize bool) (out string) {
	for i := 0; i+width <= len(input); i += width {
		for j := i; j < i+width; j++ {
			cell := input[j]
			if cell == 0 {
				out += zeroChar
			} else {
				char := Byte2ConsoleChar(cell)
				colour := cell/byte(len(ConsConsoleSafeChars)) + 1
				if colourize {
					out += fmt.Sprintf("\033[3%dm%v\033[0m", colour, string(char))
				} else {
					out += fmt.Sprintf("%v", string(char))

				}
			}
		}
		out += "\n"
	}

	return
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
