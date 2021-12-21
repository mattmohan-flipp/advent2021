package main

import (
	"fmt"
	"math"
	"strconv"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type PuzzleInput []byte

func byteArrayToUint64(b []byte) (out uint64) {
	out, _ = strconv.ParseUint(string(b), 2, 64)
	return out
}

func ProcessInput(input []string) PuzzleInput {
	line := input[0]
	str := make([]byte, 0, len(line)*4)
	for ch := 0; ch < len(input[0]); ch++ {
		num, _ := strconv.ParseUint(string(line[ch]), 16, 64)
		str = append(str, []byte(fmt.Sprintf("%04b", num))...)
	}
	return str
}

type Packet struct {
	body       []byte
	subPackets []Packet
	version    uint64
	typ        uint64
	literal    uint64
	endBit     uint
}

func decodePacket(s []byte) (p Packet) {
	p = Packet{}

	if len(s) < 11 {
		panic("Ouch")
	}

	p.body = s
	p.version = p.getBitsAtAsInt(0, 3)
	p.typ = p.getBitsAtAsInt(3, 3)

	if p.isLiteral() {
		i := uint(6)
		acc := ""
		for end := false; !end && i+4 < uint(len(p.body)); i += 5 {
			b := p.getBitsAt(i, 5)
			if b[0] == '0' {
				end = true
			}
			acc += string(b[1:])
		}
		p.endBit = i
		cur, _ := strconv.ParseUint(acc, 2, 64)
		p.literal = cur
	} else {
		p.subPackets = make([]Packet, 0, len(p.body)/11) // Min packet size is 11
		lBit := p.getBitsAt(6, 1)[0]
		if lBit == '0' {
			length := p.getBitsAtAsInt(7, 15)
			p.endBit = 22 + uint(length)
			for idx := uint(22); idx < p.endBit; {
				subPacket := decodePacket(p.body[idx:])
				idx += subPacket.endBit
				p.subPackets = append(p.subPackets, subPacket)
			}
		} else {
			packets := p.getBitsAtAsInt(7, 11)
			idx := uint(18)

			for i := uint64(0); i < packets; i++ {
				subPacket := decodePacket(p.body[idx:])
				p.subPackets = append(p.subPackets, subPacket)
				idx += subPacket.endBit
			}
			p.endBit = idx
		}
		p.literal = p.solve(0)
	}

	return
}

func (p Packet) getBitsAt(bitOffset uint, bitLength uint) []byte {
	return p.body[bitOffset : bitOffset+bitLength]
}

func (p Packet) getBitsAtAsInt(bitOffset uint, bitLength uint) uint64 {
	return byteArrayToUint64(p.getBitsAt(bitOffset, bitLength))
}

func (p Packet) isLiteral() bool {
	return p.typ == 4
}

func (p Packet) sumVersions() uint64 {
	acc := p.version
	for _, pckt := range p.subPackets {
		acc += pckt.sumVersions()
	}
	return acc
}
func (p Packet) solve(depth int) uint64 {
	switch p.typ {
	case 0:
		acc := uint64(0)
		for _, i := range p.subPackets {
			acc += i.literal
		}
		return acc
	case 1:
		acc := uint64(1)
		for _, i := range p.subPackets {
			acc *= i.literal
		}
		return acc
	case 2:
		acc := uint64(math.MaxUint64)
		for _, i := range p.subPackets {
			acc = ah.MinUint64(i.literal, acc)
		}
		return acc
	case 3:
		acc := uint64(5)
		for _, i := range p.subPackets {
			acc = ah.MaxUint64(i.literal, acc)
		}
		return acc
	case 4:
		return p.literal
	case 5:
		if p.subPackets[0].literal > p.subPackets[1].literal {
			return 1
		} else {
			return 0
		}
	case 6:
		if p.subPackets[0].literal < p.subPackets[1].literal {
			return 1
		} else {
			return 0
		}
	case 7:
		if p.subPackets[0].literal == p.subPackets[1].literal {
			return 1
		} else {
			return 0
		}
	}
	return 0
}
func SolveA(input PuzzleInput) string {
	pkt := decodePacket(input)
	return fmt.Sprint(pkt.sumVersions())
}

func SolveB(input PuzzleInput) string {

	// This is apparently subtly wrong, but at this point I'm too broken to keep going... :(
	pkt := decodePacket(input)

	return fmt.Sprint(pkt.literal)
}

func main() {
	inputLines := ah.ReadInput("day16/input.txt")

	fmt.Println("day16 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day16 Part b: ", SolveB(ProcessInput(inputLines)))
}
