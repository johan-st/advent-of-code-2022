package search

import "fmt"

type graph map[int]node

type node struct {
	neighbours  []int
	distance    int
	predecessor int
}

func (g graph) getNeighboursOf(id int) ([]int, error) {
	if node, ok := g[id]; ok {
		return node.neighbours, nil
	} else {
		return []int{}, fmt.Errorf("id %d is not in this graph", id)
	}

}
