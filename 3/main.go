package main

import (
	"fmt"
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

	// PART 2
	// split input in teams of 3
	teams := getTeams(splitInput)

	// find the common item among them
	badges := []rune{}
	for _, t := range teams {
		badges = append(badges, idTeam(t))
	}

	// score these items
	score2 := 0
	for _, b := range badges {
		fmt.Println(string(b), scoreItem(b))
		score2 += scoreItem(b)
	}

	// print score
	fmt.Println("teams aggregate priority: ", score2)
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

// PART 2
func getTeams(packs []string) [][]string {
	teams := [][]string{}
	for i, pack := range packs {
		num := int(i / 3)
		if len(teams) < num+1 {
			teams = append(teams, []string{})
		}
		teams[num] = append(teams[num], pack)
	}
	return teams
}

func idTeam(team []string) rune {
	if len(team) != 3 {
		str := fmt.Sprintf("found %d teams. Must be 3.", len(team))
		panic(str)
	}
	p1 := team[0]
	p2 := team[1]
	p3 := team[2]
	for _, r := range p1 {
		if strings.ContainsRune(p2, r) && strings.ContainsRune(p3, r) {
			fmt.Println(string(r))
			return r
		}
	}
	panic("nope")
}
