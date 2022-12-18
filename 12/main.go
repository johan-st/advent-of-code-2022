package main

import (
	"fmt"
	"strings"

	// u "github.com/johan-st/advent-of-code-2022/util"
	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	heightMap := mapFromString(u.Load("sample.txt"))
	// fmt.Printf("main %p\n", &heightMap)
	fmt.Println(heightMap)
	path := search(&heightMap)
	fmt.Println(path)
}
func search(h *heightMap) []pos {
	// fmt.Printf("search %p\n", h)
	q := que{}
	fmt.Printf("que search %p\n", &q)
	q.enque(h.start)
	prev := h.start
	(*h.nodes)[prev.row][prev.col].dist = -1
	for !q.empty() {
		dist := (*h.nodes)[prev.row][prev.col].dist + 1
		fmt.Printf("\ndist:%d\n", dist)
		fmt.Printf("que:%d\n", len(q))
		fmt.Printf("before visit (ext):%p\n", h.nodes)
		visit(h, prev, dist, q.peek())
		fmt.Printf("after visit (ext):%p\n", h.nodes)
		prev = q.peek()
		explore(&q, h)
		fmt.Printf("que post exp:%d\n", len(q))

		// TODO: debug
		if dist > 1 {
			fmt.Println("failed to find end within reasonable distance")
			fmt.Println(h)
			break
		}
	}
	return backtrack(*h)
}

func visit(h *heightMap, prev pos, dist int, p pos) {
	// fmt.Printf("visit %p\n", h)
	fmt.Println("visiting:", p, " prev:", prev)
	fmt.Printf("before visit (int): %p\n", h.nodes)
	(*h.nodes)[p.col][p.row].dist = dist
	(*h.nodes)[p.col][p.row].predecessor = prev
	(*h.nodes)[p.col][p.row].visited = true
	fmt.Printf("after visit (int): %p\n", h.nodes)

}

func explore(q *que, h *heightMap) {
	// fmt.Printf("que explore %p\n", q)
	fmt.Printf("explore:%v\n", q.peek())
	p := q.deque()
	nps := validNeighbours(*h, p)
	fmt.Printf("validNeighbours:%v\n", nps)
	for _, np := range nps {
		areDifferent := np.col != p.col || np.row != p.row
		// fmt.Println("neighbour:", np.row, np.col, "prev:", p.row, p.col, "areDifferent:", areDifferent)
		if areDifferent {
			(*h.nodes)[np.row][np.col].predecessor = p
			(*h.nodes)[np.row][np.col].dist = (*h.nodes)[p.row][p.col].dist + 1
			q.enque(np)
			if np.col == h.end.col && np.row == h.end.row {
				fmt.Printf("--FOUND END--")
				return
			}
		}
	}
}

func validNeighbours(h heightMap, p pos) []pos {
	ps := []pos{}
	n := (*h.nodes)[p.row][p.col]

	// north
	r := p.row - 1
	c := p.col
	if r >= 0 && !(*h.nodes)[r][c].visited {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}

	}

	// south
	r = p.row + 1
	c = p.col
	if r < len((*h.nodes))-1 && !(*h.nodes)[r][c].visited {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}
	}

	// west
	r = p.row
	c = p.col - 1
	if c >= 0 && !(*h.nodes)[r][c].visited {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}
	}

	// east
	r = p.row
	c = p.col + 1
	if c < len((*h.nodes)[r])-1 && !(*h.nodes)[r][c].visited {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}
	}
	return ps
}

func backtrack(h heightMap) []pos {
	return []pos{}
}

type que []pos

func (q *que) enque(p pos) {
	*q = append(*q, p)
}
func (q *que) deque() pos {
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}
func (q *que) empty() bool {
	return len(*q) == 0
}
func (q *que) peek() pos {
	return (*q)[0]
}

func mapFromString(str string) heightMap {
	rows := strings.Split(str, "\r\n")
	hm := heightMap{nodes: &[][]node{}}
	for i, row := range rows {
		newRow := []node{}
		for j, r := range row {
			if r == 'S' {
				hm.start.row = i
				hm.start.col = j
			}
			if r == 'E' {
				hm.end.row = i
				hm.end.col = j
			}
			newRow = append(newRow, newNode(runeToHeight(r), pos{i, j}))
		}
		*hm.nodes = append(*hm.nodes, newRow)
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

func newNode(h int, p pos) node {
	return node{height: h, visited: false, position: p}
}

func heightToRune(i int) rune {
	return rune(i + 97)
}

type node struct {
	position    pos
	height      int
	predecessor pos
	dist        int
	visited     bool
}

type heightMap struct {
	nodes *[][]node
	start pos
	end   pos
}

func (m heightMap) String() string {
	str := fmt.Sprintf("%v -> %v\n", m.start, m.end)
	for i, row := range *m.nodes {
		for j, n := range row {
			if i == m.start.row && j == m.start.col {
				str += "S"
			} else if i == m.end.row && j == m.end.col {
				str += "E"
			} else {
				str += string(heightToRune(n.height))
			}
		}
		str += "\n"
	}

	return str
}

type pos struct {
	row, col int
}

func (p pos) String() string {
	return fmt.Sprintf("[%d,%d]", p.row, p.col)
}
