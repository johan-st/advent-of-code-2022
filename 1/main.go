package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("running...")
	input := loadInput()
	splitInput := splitInput(input)
	calories := []int{}
	for _, value := range splitInput {
		calories = append(calories, sumString(value))
	}

	sort.Ints(calories)
	fmt.Println("Three largest are: ", calories[len(calories)-1], calories[len(calories)-2], calories[len(calories)-3])
	fmt.Println("Total: ", calories[len(calories)-1]+calories[len(calories)-2]+calories[len(calories)-3])
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
	return strings.Split(input, "\r\n\r\n")
}

func sumString(input string) int {
	if input == "" {
		return 0
	}
	split := strings.Split(input, "\r\n")
	sum := 0
	for _, value := range split {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		sum += num
	}
	return sum
}

func getMax(input []int) int {
	if len(input) == 0 {
		return 0
	}
	max := input[0]
	for _, value := range input {
		if value > max {
			max = value
		}
	}
	return max
}
