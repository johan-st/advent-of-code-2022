package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Load input
	input := load("input.txt")
	pairs := parsePairs(input)
	numContained := 0
	for i, p := range pairs {
		contained := zonesFullyContained(p.firstZones, p.secondZones)
		fmt.Println(i, ":", p, " - ", contained)
		if contained {
			numContained++
		}
	}
	fmt.Println("Number of pairs fully contained:", numContained)
}

func load(filename string) string {
	// Load input
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}

// PAIRS
type pair struct {
	firstZones  zones
	secondZones zones
}
type zones struct {
	from zone
	to   zone
}
type zone int

func parsePairs(input string) []pair {
	splitInput := strings.Split(input, "\r\n")
	pairs := []pair{}
	for _, p := range splitInput {
		pairs = append(pairs, parsePair(p))
	}
	return pairs
}

func parsePair(pairString string) pair {
	p := pair{}
	splitPair := strings.Split(pairString, ",")
	p.firstZones = parseZones(splitPair[0])
	p.secondZones = parseZones(splitPair[1])
	return p
}

func parseZones(input string) zones {
	zones := zones{}
	splitZones := strings.Split(input, "-")
	min, err := strconv.Atoi(splitZones[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(splitZones[1])
	if err != nil {
		panic(err)
	}
	zones.from = zone(min)
	zones.to = zone(max)

	return zones
}

func zonesFullyContained(z1 zones, z2 zones) bool {
	if z1.from >= z2.from && z1.to <= z2.to {
		return true
	}
	if z2.from >= z1.from && z2.to <= z1.to {
		return true
	}
	return false
}
