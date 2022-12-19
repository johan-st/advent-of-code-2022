package search

import (
	"fmt"
)

type Graph map[int]node

type node struct {
	id          int
	neighbours  []int
	predecessor int
}

// edge {from, to} if used as a directed edge (as opposed to undirected)
type Edge [2]int

func (g Graph) AddNodes(ids []int) error {
	for _, id := range ids {
		err := g.addNode(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g Graph) addNode(id int) error {
	if id < 1 {
		return fmt.Errorf("id must be greater than 0. id given was %d", id)

	}
	if _, ok := g[id]; ok {
		return fmt.Errorf("node with id %d already exists", id)
	}
	g[id] = node{id: id}
	return nil
}

func (g Graph) addEdge(e Edge) error {
	if e[0] < 1 && e[1] < 1 {
		return fmt.Errorf("could not add edge between node %d and %d. Ids must be greater than 0", e[0], e[1])
	}

	var n1, n2 node
	if node, ok := g[e[0]]; !ok {
		return fmt.Errorf("could not add edge between node %d and %d. %d does not exist", e[0], e[1], e[0])
	} else {
		n1 = node
	}
	if node, ok := g[e[1]]; !ok {
		return fmt.Errorf("could not add edge between node %d and %d. %d does not exist", e[0], e[1], e[1])
	} else {
		n2 = node
	}

	n1.neighbours = append(g[e[0]].neighbours, e[1])
	n2.neighbours = append(g[e[1]].neighbours, e[0])
	g[e[0]] = n1
	g[e[1]] = n2
	return nil
}

func (g Graph) AddEdges(es []Edge) error {
	for _, e := range es {
		err := g.addEdge(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g Graph) GetNeighboursFor(id int) []int {
	node := g[id]
	return node.neighbours
}

func (g Graph) addEdge_directional(e Edge) error {
	if e[0] < 1 && e[1] < 1 {
		return fmt.Errorf("could not add edge between node %d and %d. Ids must be greater than 0", e[0], e[1])
	}

	var n1 node
	if node, ok := g[e[0]]; !ok {
		return fmt.Errorf("could not add edge between node %d and %d. %d does not exist", e[0], e[1], e[0])
	} else {
		n1 = node
	}

	n1.neighbours = append(g[e[0]].neighbours, e[1])
	g[e[0]] = n1
	return nil
}

func (g Graph) AddEdges_directional(es []Edge) error {
	for _, e := range es {
		err := g.addEdge_directional(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g Graph) BredthFirst(start int, end int) []int {
	q := que{}
	q.enque(g[start])
	for !q.empty() {
		current := q.deque()
		for _, n := range current.neighbours {
			next := g[n]
			if n == end {
				// TODO:
				node := g[n]
				node.predecessor = current.id
				g[n] = node
				q = que{}
				break
			} else if next.predecessor == 0 { //predecessor should only be 0 for not visited nodes
				next.predecessor = current.id
				g[n] = next
				q.enque(next)
			}
		}
	}

	path := []int{end}
	if g[end].predecessor == 0 {
		return []int{}
	}
	last := end
	done := false
	for !done {
		id := g[last].predecessor
		last = id
		path = append(path, g[last].id)
		if id == start {
			done = true
		}
		if id < 1 {
			panic("\n\nsomething went very wrong here. This might be related to predecessors or an unsolvable path\n")
		}
	}
	path = reverse(path)
	return path
}

func reverse(ints []int) []int {
	rev := []int{}
	for i := range ints {
		rev = append(rev, ints[len(ints)-1-i])
	}
	return rev
}

// Que
type que []node

func (q *que) enque(n node) {
	*q = append(*q, n)
}
func (q *que) deque() node {
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}
func (q *que) empty() bool {
	return len(*q) == 0
}
func (q *que) peek() node {
	return (*q)[0]
}
