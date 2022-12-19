package main

import (
	"fmt"
	"strings"

	// u "github.com/johan-st/advent-of-code-2022/util"
	u "github.com/johan-st/advent-of-code-2022/util"
	// search "github.com/johan-st/advent-of-code-2022/util/search"
	s "github.com/johan-st/advent-of-code-2022/util/search"
)

func main() {
	heightMap := mapFromString(u.Load("input.txt"))
	// fmt.Println(heightMap)
	graph, start, end := graphFromMap(heightMap)
	shortestPath := graph.BredthFirst(start, end)
	if len(shortestPath) < 2 {
		fmt.Println("failed to establish path")
	} else {
		fmt.Println("shortest path is", len(shortestPath)-1, "steps")
	}

	shortestHike := len(shortestPath)
	allLowest := lowestPositions(heightMap)
	for i, id := range allLowest {
		fmt.Printf("\rCalculating (id %d) %d of %d. currently shortest is %d", id, i, len(allLowest), shortestHike-1)
		graph, _, end := graphFromMap(heightMap)
		short := graph.BredthFirst(id, end)
		if len(short) < shortestHike && len(short) > 1 {
			shortestHike = len(short)
		}
	}
	fmt.Println("\nshortest hike is ", shortestHike-1, "steps")
}

func lowestPositions(m heightMap) []int {
	low := []int{}
	for i, row := range *m.nodes {
		for j, node := range row {
			if node.height == 0 {
				id := i*len(row) + j + 1
				low = append(low, id)
			}

		}
	}
	return low
}
func graphFromMap(m heightMap) (s.Graph, int, int) {
	g := s.Graph{}

	count := 0
	edges := []s.Edge{}
	start := m.start.row*len((*m.nodes)[0]) + m.start.col + 1
	end := m.end.row*len((*m.nodes)[0]) + m.end.col + 1
	for i, row := range *m.nodes {
		for j := range row {
			count++
			ns := validNeighbours(m, pos{i, j})
			for _, n := range ns {
				idN := n.row*len(row) + n.col + 1
				edges = append(edges, s.Edge{count, idN})
			}

		}
	}

	nodes := []int{}
	for id := 1; id <= count; id++ {
		nodes = append(nodes, id)
	}

	err := g.AddNodes(nodes)
	if err != nil {
		panic(err)
	}
	err = g.AddEdges_directional(edges)
	if err != nil {
		panic(err)
	}
	return g, start, end
}

func validNeighbours(h heightMap, p pos) []pos {
	ps := []pos{}
	n := (*h.nodes)[p.row][p.col]

	// north
	r := p.row - 1
	c := p.col
	if r >= 0 {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}

	}

	// south
	r = p.row + 1
	c = p.col
	if r < len((*h.nodes)) {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}
	}

	// west
	r = p.row
	c = p.col - 1
	if c >= 0 {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}
	}

	// east
	r = p.row
	c = p.col + 1
	if c < len((*h.nodes)[r]) {
		h := (*h.nodes)[r][c].height
		if h <= n.height+1 {
			ps = append(ps, pos{r, c})
		}
	}
	return ps
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
	return node{h}
}

func heightToRune(i int) rune {
	return rune(i + 97)
}

type node struct {
	height int
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
