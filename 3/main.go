package main

import (
	"os"
	"strings"
)

func main() {
	// Load input
	input := load("input.txt")

	// Split input
	splitInput := strings.Split(input, "\r\n")

	// Find duplicates
	duplicates := findDuplicates(splitInput)

	// score duplicates
	score := 0
	for _, d := range duplicates {
		score += scoreDuplicate(d)
	}

	// print score
	println(score)
}

func load(filename string) string {
	// Load input
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func findDuplicates(in []string) []rune {
	if len(in)%2 != 0 || len(in) == 0 {
		return []rune{}
	}
	return []rune{'a'}
	// return []rune{rune(in[0][0]), rune(in[1][0])}
}

func scoreDuplicate(d rune) int {
	if d > 90 {
		return (int(d) - 96)
	} else {
		return (int(d) - 38)
	}
}
