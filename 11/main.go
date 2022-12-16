package main

import (
	"fmt"
	"sort"
	// u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	fmt.Println("MONKEY BUSSINES v.1")
	state := initialState()

	commonDivisor := 1
	for _, m := range *state {
		commonDivisor *= m.divisor
	}
	fmt.Println("commonDivisor", commonDivisor)

	steps := 10000
	for i := 0; i < steps; i++ {
		runRound(state)
	}
	sort.Sort(state)
	fmt.Println(state)
	monkeyBusiness := (*state)[len(*state)-1].numInspections * (*state)[len(*state)-2].numInspections
	fmt.Println("Monkey Business:", monkeyBusiness)
}
func initialState() *state {
	// state := sample()
	state := input()
	return &state
}

type monkey struct {
	items           []int
	operation       func(int) int
	test            func(int) bool
	divisor         int
	monkeyWhenTrue  int
	monkeyWhenFalse int
	numInspections  int
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
	m.numInspections++

	// This magic number is the multiplication of all divBy test numbers.
	// This hinders overflow whitout impacting the tests
	// sample: 96577, input: 9699690
	// m.items[0] = (m.operation(m.items[0])) % 96577
	m.items[0] = (m.operation(m.items[0])) % 9699690
	// fmt.Println(m.items)

	if m.test(m.items[0]) {
		throw(s, monkeyIndex, m.monkeyWhenTrue)
	} else {
		throw(s, monkeyIndex, m.monkeyWhenFalse)

	}
}

type state []monkey

func (a state) Len() int           { return len(a) }
func (a state) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a state) Less(i, j int) bool { return a[i].numInspections < a[j].numInspections }

func throw(s *state, from int, to int) {
	a := (*s)[from].items[0]
	(*s)[to].items = append((*s)[to].items, a)
	(*s)[from].items = (*s)[from].items[1:]

}

func (s state) String() string {
	str := fmt.Sprintf("%d monkeys in state", len(s))
	for i, monkey := range s {
		str += fmt.Sprintf("\nMonkey %d inspected items %d times.", i, monkey.numInspections)
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
	m1 := monkey{items: []int{79, 98}, operation: mul(19), test: div(23), divisor: 23, monkeyWhenTrue: 2, monkeyWhenFalse: 3}
	m2 := monkey{items: []int{54, 65, 75, 74}, operation: add(6), test: div(19), divisor: 19, monkeyWhenTrue: 2, monkeyWhenFalse: 0}
	m3 := monkey{items: []int{79, 60, 97}, operation: square, test: div(13), divisor: 13, monkeyWhenTrue: 1, monkeyWhenFalse: 3}
	m4 := monkey{items: []int{74}, operation: add(3), test: div(17), divisor: 17, monkeyWhenTrue: 0, monkeyWhenFalse: 1}

	return state{m1, m2, m3, m4}
}
func input() state {
	return state{
		monkey{[]int{75, 63}, mul(3), div(11), 11, 7, 2, 0},
		monkey{[]int{65, 79, 98, 77, 56, 54, 83, 94}, add(3), div(2), 2, 2, 0, 0},
		monkey{[]int{66}, add(5), div(5), 5, 7, 5, 0},
		monkey{[]int{51, 89, 90}, mul(19), div(7), 7, 6, 4, 0},
		monkey{[]int{75, 94, 66, 90, 77, 82, 61}, add(1), div(17), 17, 6, 1, 0},
		monkey{[]int{53, 76, 59, 92, 95}, add(2), div(19), 19, 4, 3, 0},
		monkey{[]int{81, 61, 75, 89, 70, 92}, square, div(3), 3, 0, 1, 0},
		monkey{[]int{81, 86, 62, 87}, add(8), div(13), 13, 3, 5, 0},
	}
}

// helpers

func mul(a int) func(int) int {
	return func(b int) int { return a * b }
}

func add(a int) func(int) int {
	return func(b int) int { return a + b }
}

func div(div int) func(int) bool {
	return func(a int) bool { return a%div == 0 }
}

func square(a int) int {
	return a * a
}
