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
	fmt.Println("Part 2, most scenic tree: ", mostScenicTree(f))

}

func mostScenicTree(f forest) int {
	var max int
	for i, row := range f {
		for j := range row {
			if f.scenicValueAt(i, j) > max {
				max = f[i][j].height
			}
		}
	}
	return max
}

type tree struct {
	height      int
	scenicValue int
	scenicLeft  int
	scenicRight int
	scenicUp    int
	scenicDown  int
}
type forest [][]tree

func (f forest) scenicValueAt(i, j int) int {
	return 1
}
func (f forest) numVisible() int {
	var num int
	for i, row := range f {
		for j := range row {
			if f.isVisible(i, j) {
				num++
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
		if f[i][k].height >= f[i][j].height {
			return false
		}

	}
	return true
}
func (f forest) isVisibleRight(i, j int) bool {
	for k := j + 1; k < len(f[i]); k++ {
		if f[i][k].height >= f[i][j].height {
			return false
		}
	}
	return true
}
func (f forest) isVisibleUp(i, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if f[k][j].height >= f[i][j].height {
			return false
		}
	}
	return true
}
func (f forest) isVisibleDown(i, j int) bool {
	for k := i + 1; k < len(f); k++ {
		if f[k][j].height >= f[i][j].height {
			return false
		}
	}
	return true
}

func (f forest) String() string {
	var s string
	for i, row := range f {
		for _, col := range row {
			s += fmt.Sprintf("%d", col.height)
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
		f = append(f, []tree{})
		for _, col := range row {
			height, err := strconv.Atoi(string(col))
			if err != nil {
				panic(err)
			}
			f[i] = append(f[i], tree{height: height})
		}
		s += "\n"
	}
	return f
}
