package main

import (
	"fmt"

	ah "github.com/mattmohan/advent2021/pkg/advent_helpers"
)

type SnailTreeNode struct {
	// For tree based navigation
	LeftNode  *SnailTreeNode
	RightNode *SnailTreeNode

	// For list based navigation
	previousNode *SnailTreeNode
	nextNode     *SnailTreeNode

	// To walk up the tree if needed
	parentNode *SnailTreeNode

	Value byte
}

func (n SnailTreeNode) isLeaf() bool {
	return n.LeftNode == nil && n.RightNode == nil
}

func (n *SnailTreeNode) addChild(c *SnailTreeNode) {
	c.parentNode = n
	if n.LeftNode == nil {
		n.LeftNode = c
	} else if n.RightNode == nil {
		n.RightNode = c
	} else {
		panic("Too many children")
	}
}

func (n SnailTreeNode) Magnitude() int {
	if n.isLeaf() {
		return int(n.Value)
	}

	return n.LeftNode.Magnitude()*3 + n.RightNode.Magnitude()*2
}

func (a *SnailTreeNode) addNode(b *SnailTreeNode) *SnailTreeNode {
	newNode := new(SnailTreeNode)
	newNode.LeftNode = a
	newNode.RightNode = b
	a.parentNode = newNode
	b.parentNode = newNode

	aRight := a
	for aRight.RightNode != nil {
		aRight = aRight.RightNode
	}

	bLeft := b
	for bLeft.LeftNode != nil {
		bLeft = bLeft.LeftNode
	}

	aRight.nextNode = bLeft
	bLeft.previousNode = aRight

	return newNode
}

func (n *SnailTreeNode) explode(depth int) bool {
	if depth >= 4 && !n.isLeaf() {
		n.Value = 0

		if n.LeftNode.previousNode != nil {
			n.LeftNode.previousNode.Value += n.LeftNode.Value
			n.LeftNode.previousNode.nextNode = n
		}
		if n.RightNode.nextNode != nil {
			n.RightNode.nextNode.Value += n.RightNode.Value
			n.RightNode.nextNode.previousNode = n
		}
		n.previousNode = n.LeftNode.previousNode
		n.nextNode = n.RightNode.nextNode
		n.LeftNode = nil
		n.RightNode = nil
		return true
	}
	if n.LeftNode != nil && n.LeftNode.explode(depth+1) {
		return true
	}
	if n.RightNode != nil && n.RightNode.explode(depth+1) {
		return true
	}
	return false
}
func (n *SnailTreeNode) split() bool {
	if n.Value > 9 {
		n.LeftNode = new(SnailTreeNode)
		n.RightNode = new(SnailTreeNode)

		n.LeftNode.previousNode = n.previousNode
		n.LeftNode.nextNode = n.RightNode
		n.LeftNode.parentNode = n
		n.LeftNode.Value = n.Value / 2

		n.RightNode.nextNode = n.nextNode
		n.RightNode.previousNode = n.LeftNode
		n.RightNode.parentNode = n
		n.RightNode.Value = n.Value/2 + (n.Value % 2)

		if n.previousNode != nil {
			n.previousNode.nextNode = n.LeftNode
		}
		if n.nextNode != nil {
			n.nextNode.previousNode = n.RightNode
		}

		n.Value = 0
		return true
	}
	if n.LeftNode != nil && n.LeftNode.split() {
		return true
	}
	if n.RightNode != nil && n.RightNode.split() {
		return true
	}
	return false
}
func (n *SnailTreeNode) Reduce() bool {
	for {
		if n.LeftNode != nil && n.LeftNode.explode(1) {
			continue
		}
		if n.RightNode != nil && n.RightNode.explode(1) {
			continue
		}
		if n.LeftNode != nil && n.LeftNode.split() {
			continue
		}
		if n.RightNode != nil && n.RightNode.split() {
			continue
		}
		return false
	}
}

func (n SnailTreeNode) String() string {
	if n.isLeaf() {
		return fmt.Sprint(n.Value)
	}
	return fmt.Sprintf("[%v,%v]", n.LeftNode.String(), n.RightNode.String())
}

type PuzzleInput []string

func ProcessInput(input []string) PuzzleInput {
	return input
}

func ParseSnailNumber(input string) *SnailTreeNode {
	stack := make([]*SnailTreeNode, 0, len(input)/5)
	rootNode := SnailTreeNode{}
	stack = append(stack, &rootNode)

	var previousNumNode *SnailTreeNode

	for i := 1; i < len(input); i++ {
		cur := stack[len(stack)-1]
		if input[i] == '[' {
			newNode := new(SnailTreeNode)
			stack = append(stack, newNode)
			cur.addChild(newNode)
		} else if input[i] == ']' {
			stack[len(stack)-1] = nil
			stack = stack[:len(stack)-1]
		} else if input[i] >= '0' && input[i] <= '9' {
			newNode := new(SnailTreeNode)

			if previousNumNode != nil {
				previousNumNode.nextNode = newNode
			}
			newNode.previousNode = previousNumNode
			previousNumNode = newNode

			newNode.Value = input[i] - '0'
			cur.addChild(newNode)
		}
	}
	return &rootNode
}

func SolveA(input PuzzleInput) string {
	num := ParseSnailNumber(input[0])

	for i := 1; i < len(input); i++ {
		cur := ParseSnailNumber(input[i])
		num = num.addNode(cur)
		num.Reduce()
	}

	return fmt.Sprint(num.Magnitude())
}

func SolveB(input PuzzleInput) string {
	max := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			first := ParseSnailNumber(input[i])
			second := ParseSnailNumber(input[j])
			first = first.addNode(second)
			first.Reduce()
			max = ah.MaxInt(max, first.Magnitude())
		}
	}

	return fmt.Sprint(max)
}

func main() {
	inputLines := ah.ReadInput("day18/input.txt")

	fmt.Println("day18 Part a: ", SolveA(ProcessInput(inputLines)))
	fmt.Println("day18 Part b: ", SolveB(ProcessInput(inputLines)))
}
