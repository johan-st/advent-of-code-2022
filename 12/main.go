package main

import (
	"fmt"
	"strings"

	// u "github.com/johan-st/advent-of-code-2022/util"
	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	heightMap := mapFromString(u.Load("sample.txt"))
	fmt.Println(heightMap)
}

func mapFromString(str string) heightMap {
	rows := strings.Split(str, "\r\n")
	hm := heightMap{}
	for i, row := range rows {
		newRow := []int{}
		for j, r := range row {
			if r == 'S' {
				hm.start.row = i
				hm.start.col = j
			}
			if r == 'E' {
				hm.end.row = i
				hm.end.col = j
			}
			newRow = append(newRow, runeToHeight(r))
		}
		hm.heights = append(hm.heights, newRow)
	}
	return hm
}
func runeToHeight(r rune) int {
	if r == 'S' {
		return 0
	}
	if r == 'E' {
		return 25
	}
	return int(r - 97)
}

func heightToRune(i int) rune {
	return rune(i + 97)
}

type heightMap struct {
	heights [][]int
	start   pos
	end     pos
}

func (m heightMap) String() string {
	str := fmt.Sprintf("%v -> %v\n", m.start, m.end)
	for i, row := range m.heights {
		for j, h := range row {
			if i == m.start.row && j == m.start.col {
				str += "S"
			} else if i == m.end.row && j == m.end.col {
				str += "E"
			} else {
				str += string(heightToRune(h))
			}
		}
		str += "\n"
	}

	return str
}

type pos struct {
	row, col int
}
