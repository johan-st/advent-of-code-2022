package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := loadInput()
	splitInput := splitInput(input)
	total := 0
	for _, v := range splitInput {
		total += scoreGame(v)
	}
	fmt.Println("Part 1: ", total)
	totalPart2 := 0
	for _, v := range splitInput {
		totalPart2 += scoreGamePart2(v)
	}
	fmt.Println("Part 2: ", totalPart2)
}

// read input from file
func loadInput() string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}

func splitInput(input string) []string {
	if input == "" {
		return nil
	}
	return strings.Split(input, "\r\n")
}

func scoreGame(game string) int {
	if game == "A X" {
		return 4
	} else if game == "A Y" {
		return 8
	} else if game == "A Z" {
		return 3
	} else if game == "B X" {
		return 1
	} else if game == "B Y" {
		return 5
	} else if game == "B Z" {
		return 9
	} else if game == "C X" {
		return 7
	} else if game == "C Y" {
		return 2
	} else if game == "C Z" {
		return 6
	} else {
		return 0
	}
}

func scoreGamePart2(game string) int {
	if game == "A X" {
		return 3
	} else if game == "A Y" {
		return 4
	} else if game == "A Z" {
		return 8
	} else if game == "B X" {
		return 1
	} else if game == "B Y" {
		return 5
	} else if game == "B Z" {
		return 9
	} else if game == "C X" {
		return 2
	} else if game == "C Y" {
		return 6
	} else if game == "C Z" {
		return 7
	} else {
		return 0
	}
}
