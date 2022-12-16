package main

import (
	"fmt"

	// u "github.com/johan-st/advent-of-code-2022/util"
	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	fmt.Println("MONKEY BUSSINES v.1")
	state := stateFromString(u.Load("sample.txt"))
	steps := 20
	fmt.Println(0, state)
	for i := 0; i < steps; i++ {
		runRound(state)
		fmt.Println(i+1, state)
	}
}

type monkey struct {
	items           []int
	operation       func(int) int
	test            func(int) bool
	monkeyWhenTrue  int
	monkeyWhenFalse int
}

func runRoundMonkey(s *state, monkeyIndex int) {
	monkey := (*s)[monkeyIndex]
	for range monkey.items {
		inspectAndAct(s, monkeyIndex)
	}
}

func inspectAndAct(s *state, monkeyIndex int) {
	// fmt.Println(m.items)
	m := &(*s)[monkeyIndex]
	m.items[0] = (m.operation(m.items[0])) / 3
	// fmt.Println(m.items)

	if m.test(m.items[0]) {
		throw(s, monkeyIndex, m.monkeyWhenTrue)
	} else {
		throw(s, monkeyIndex, m.monkeyWhenFalse)

	}
}

type state []monkey

func stateFromString(str string) *state {
	// TODO: read from file
	state := sample()
	return &state
}

func throw(s *state, from int, to int) {
	a := (*s)[from].items[0]
	(*s)[to].items = append((*s)[to].items, a)
	(*s)[from].items = (*s)[from].items[1:]

}

func (s state) String() string {
	str := fmt.Sprintf("%d monkeys in state", len(s))
	for i, monkey := range s {
		str += fmt.Sprintf("\nMonkey %d: %v", i, monkey.items)
	}
	return str
}

func runRound(s *state) {
	for i := 0; i < len(*s); i++ {
		runRoundMonkey(s, i)
	}
}

// SAMPLE MONKEY
func sample() state {
	m1 := monkey{items: []int{79, 98}, operation: multiply(19), test: divisible(23), monkeyWhenTrue: 2, monkeyWhenFalse: 3}
	m2 := monkey{items: []int{54, 65, 75, 74}, operation: add(6), test: divisible(19), monkeyWhenTrue: 2, monkeyWhenFalse: 0}
	m3 := monkey{items: []int{79, 60, 97}, operation: square, test: divisible(13), monkeyWhenTrue: 1, monkeyWhenFalse: 3}
	m4 := monkey{items: []int{74}, operation: add(3), test: divisible(17), monkeyWhenTrue: 0, monkeyWhenFalse: 1}

	return state{m1, m2, m3, m4}
}

// helpers

func multiply(a int) func(int) int {
	return func(b int) int { return a * b }
}

func add(a int) func(int) int {
	return func(b int) int { return a + b }
}

func divisible(div int) func(int) bool {
	return func(a int) bool { return a%div == 0 }
}

func square(a int) int {
	return a * a
}
