package main

import (
	"fmt"

	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	input := u.Load("input.txt")

	pos := findMarker(4, input)
	pos2 := findMarker(14, input)

	fmt.Printf("marker 1: %d\n", pos)
	fmt.Printf("marker 2: %d\n", pos2)
}

func findMarker(size int, input string) int {
	// last few characters. determined by size
	last := make([]rune, size)
	for i := 0; i < size; i++ {
		last[i] = rune(input[i])
	}
	for i := size; i < len(input); i++ {
		last[i%size] = rune(input[i])
		if allUniqueRunes(last) {
			fmt.Println(string(last))
			return i + 1
		}
	}
	return 0
}

func allUniqueRunes(r []rune) bool {
	seen := make(map[rune]bool)
	for _, v := range r {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}
