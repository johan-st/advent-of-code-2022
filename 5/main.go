package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/johan-st/advent-of-code-2022/util"
	s "github.com/johan-st/advent-of-code-2022/util/stack"
)

func main() {
	input := u.Load("input.txt")
	splitInput := strings.Split(input, "\r\n\r\n")
	initialStateStr := splitInput[0]
	movesStr := splitInput[1]
	state := cargo{}.fromString(initialStateStr, 9)
	moves := moves{}.fromString(movesStr)
	state.applyMoves(moves)
	// fmt.Printf("MOVES:\n%s", moves.toString())
	fmt.Printf("TOP: %s", state.top())
	// fmt.Printf("CARGO:\n%s", state.toString())
}

// Move, from a stack to another. itterations is the number of tines the move is to be performed.
type move struct {
	from, to, itterations int
}
type moves []move

func (m moves) toString() string {
	str := ""
	for _, move := range m {

		str += fmt.Sprintf("move %d from %d to %d\n", move.itterations, move.from, move.to)
	}
	return str
}

func (m moves) fromString(str string) moves {
	rows := strings.Split(str, "\r\n")
	for _, row := range rows {
		if row != "" {
			m = append(m, parseMove(row))
		}
	}
	return m
}

func parseMove(str string) move {
	m := move{}
	split := strings.Split(str, " ")
	for i := 0; i < len(split); i++ {

	}
	itt, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	from, err := strconv.Atoi(split[3])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(split[5])
	if err != nil {
		panic(err)
	}
	m.itterations = itt
	m.from = from
	m.to = to

	return m
}

// CARGO is a number of stacks
type cargo []*s.Stack

func (c cargo) applyMoves(m moves) {
	for _, move := range m {
		c.applyMove(move)
	}
}
func (c cargo) applyMove(m move) {
	for i := 0; i < m.itterations; i++ {
		c[m.to-1].Push(c[m.from-1].Pop())
	}
}

func (c cargo) fromString(initialStateStr string, numSpaces int) cargo {
	temp := cargo{}
	for i := 0; i < numSpaces; i++ {
		temp = append(temp, s.New())
	}
	rows := strings.Split(initialStateStr, "\r\n")
	for i, row := range rows {
		fmt.Println(i, ":", row)
		if i != len(rows)-1 {
			for j := 0; j < numSpaces; j++ {
				r := row[j*4+1 : j*4+2]
				if r != " " {
					temp[j%numSpaces].Push(r)
				}
			}
		}
	}

	for i := 0; i < numSpaces; i++ {
		c = append(c, s.New())
		addStackInReverse(c[i], temp[i])
	}
	fmt.Println(temp.toString())
	return c
}
func addStackInReverse(dest *s.Stack, src *s.Stack) {
	for !src.Empty() {
		dest.Push(src.Pop())
	}
}

// NOTE: This is destructive!
// Returns a string representation of the cargo state.
func (c cargo) toString() string {
	topHeight := 0
	tmpC := cargo(c)
	for _, b := range c {
		if b.Size() > topHeight {
			topHeight = b.Size()
		}
	}
	str := ""
	for l := topHeight; l > 0; l-- {
		for _, b := range tmpC {
			if b.Size() >= l {
				str += fmt.Sprintf("[%v] ", b.Pop())
			} else {
				str += "    "
			}
		}
		str += "\n"
	}
	return str
}

func (c cargo) top() string {
	str := ""
	for _, b := range c {
		if b.Empty() {
			str += "_"
		} else {
			str += fmt.Sprintf("%s", b.Top())
		}
	}
	str += "\n"
	return str
}

// CARGO END
