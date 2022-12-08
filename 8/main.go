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
	fmt.Println("Part 2, highest scenic value: ", mostScenicTree(f))

}

func mostScenicTree(f forest) int {
	mostScenicTree := [2]int{0, 0}
	for i, row := range f {
		for j := range row {
			if f.scenicValueAt(i, j) > f.scenicValueAt(mostScenicTree[0], mostScenicTree[1]) {
				mostScenicTree = [2]int{i, j}
			}
		}
	}
	return f.scenicValueAt(mostScenicTree[0], mostScenicTree[1])
}

type tree struct {
	height      int
	scenicValue int
}
type forest [][]tree

func (f forest) scenicValueAt(i, j int) int {
	if f[i][j].scenicValue > 0 {
		return f[i][j].scenicValue
	}
	f[i][j].scenicValue = f.scenicLeft(i, j) * f.scenicRight(i, j) * f.scenicUp(i, j) * f.scenicDown(i, j)

	return f[i][j].scenicValue
}

func (f forest) scenicLeft(i, j int) int {
	var num int
	for k := j - 1; k >= 0; k-- {
		if f[i][k].height >= f[i][j].height {
			num++
			return num
		}
		num++
	}
	return num
}
func (f forest) scenicRight(i, j int) int {
	var num int
	for k := j + 1; k < len(f[i]); k++ {
		if f[i][k].height >= f[i][j].height {
			num++
			return num
		}
		num++
	}
	return num
}

func (f forest) scenicUp(i, j int) int {
	var num int
	for k := i - 1; k >= 0; k-- {
		if f[k][j].height >= f[i][j].height {
			num++
			return num
		}
		num++
	}
	return num
}
func (f forest) scenicDown(i, j int) int {
	var num int
	for k := i + 1; k < len(f); k++ {
		if f[k][j].height >= f[i][j].height {
			num++
			return num
		}
		num++
	}
	return num
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
