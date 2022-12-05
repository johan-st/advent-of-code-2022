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
	duplicates := ""
	for _, s := range splitInput {
		duplicates += string(findDuplicates(s))
	}

	// score duplicates
	score := 0
	for _, d := range duplicates {
		score += scoreItem(d)
	}

	// print score
	println("total priority: ", score)

}

func load(filename string) string {
	// Load input
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func findDuplicates(in string) rune {
	if len(in)%2 != 0 || len(in) == 0 {
		return 0
	}
	half := int(len(in) / 2)
	h := in[:half]
	t := in[half:]
	for _, r := range h {
		if strings.ContainsRune(t, r) {
			return r
		}
	}
	return 0
}

func scoreItem(d rune) int {
	if d > 90 {
		return (int(d) - 96)
	} else {
		return (int(d) - 38)
	}
}
