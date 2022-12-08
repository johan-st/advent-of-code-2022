package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	// Read input
	f := forest{}.fromString(u.Load("input.txt"))
	// fmt.Println(f)

	// part 1
	fmt.Println("Part 1, visible trees: ", f.numVisible(), "out of:", len(f)*len(f[0]), "(", len(f), "x", len(f[0]), ")")
	// part 2

}

type forest [][]int

func (f forest) numVisible() int {
	var num int
	for i, row := range f {
		for j := range row {
			if f.isVisible(i, j) {
				num++
			} else {
				fmt.Printf("[%d,%d]\n", i, j)
			}
		}
	}
	return num
}

func (f forest) isVisible(i, j int) bool {
	rows := len(f)
	cols := len(f[0])
	if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
		return true
	}
	if f.isVisibleLeft(i, j) || f.isVisibleRight(i, j) || f.isVisibleUp(i, j) || f.isVisibleDown(i, j) {
		return true
	}
	// check row-left
	// res := false

	return false
}
func (f forest) isVisibleLeft(i, j int) bool {
	for k := j - 1; k >= 0; k-- {
		if f[i][k] >= f[i][j] {
			return false
		}

	}
	return true
}
func (f forest) isVisibleRight(i, j int) bool {
	for k := j + 1; k < len(f[i]); k++ {
		if f[i][k] >= f[i][j] {
			return false
		}
	}
	return true
}
func (f forest) isVisibleUp(i, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if f[k][j] >= f[i][j] {
			return false
		}
	}
	return true
}
func (f forest) isVisibleDown(i, j int) bool {
	for k := i + 1; k < len(f); k++ {
		if f[k][j] >= f[i][j] {
			return false
		}
	}
	return true
}

func (f forest) String() string {
	var s string
	for i, row := range f {
		for _, col := range row {
			s += fmt.Sprintf("%d", col)
		}
		if i < len(f)-1 {
			s += "\r\n"
		}
	}
	return s
}

func (f forest) fromString(s string) forest {
	str := strings.Split(s, "\r\n")
	for i, row := range str {
		f = append(f, []int{})
		for _, col := range row {
			height, err := strconv.Atoi(string(col))
			if err != nil {
				panic(err)
			}
			f[i] = append(f[i], height)
		}
		s += "\n"
	}
	return f
}
