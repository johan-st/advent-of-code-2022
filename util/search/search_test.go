package search_test

import (
	"fmt"
	"reflect"
	"testing"

	s "github.com/johan-st/advent-of-code-2022/util/search"
)

func Test_Graph_AddNodes(t *testing.T) {
	ids := []int{1, 2, 3, 4, 50}
	g := s.Graph{}
	err := g.AddNodes(ids)
	if err != nil {
		fmt.Println(err)
	}
	for _, id := range ids {
		if _, ok := g[id]; !ok {
			fmt.Printf("failed to retrieve nodes added")
			t.Fail()
		}
	}
}

func Test_Graph_AddEdges(t *testing.T) {
	g := s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4, 5, 6, 99})
	es := []s.Edge{
		{1, 2}, {1, 3}, {1, 4},
		{4, 99}, {4, 5}, {4, 6},
		{99, 1}, {99, 2}, {99, 3},
	}
	g.AddEdges(es)
	for _, e := range es {
		if !checkEdge(g, e) {
			fmt.Printf("check on edges failed on edge%v\n", e)
			t.Fail()
		}
	}
}

func Test_Graph_AddEdges_directional(t *testing.T) {
	g := s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4})
	es := []s.Edge{{1, 2}, {2, 3}, {3, 4}, {4, 1}}
	g.AddEdges_directional(es)
	for _, e := range es {
		if !checkEdge_directional(g, e) {
			fmt.Printf("check on edges failed on edge%v\n", e)
			t.Fail()
		}
	}
}

func Test_Graph_BredthFirst(t *testing.T) {

	g := s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4, 5, 6})
	g.AddEdges([]s.Edge{{1, 2}, {1, 3}, {2, 3}, {3, 4}, {3, 5}, {5, 6}})
	path := g.BredthFirst(1, 2)
	want := []int{1, 2}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("wanted: %v\ngot: %v\n", want, path)
	}

	g = s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4, 5, 6})
	g.AddEdges([]s.Edge{{1, 2}, {1, 3}, {2, 3}, {3, 4}, {3, 5}, {5, 6}})
	path = g.BredthFirst(1, 4)
	want = []int{1, 3, 4}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("wanted: %v\ngot: %v\n", want, path)
	}

	g = s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4, 5, 6})
	g.AddEdges([]s.Edge{{1, 2}, {1, 3}, {2, 3}, {3, 4}, {3, 5}, {5, 6}})
	path = g.BredthFirst(5, 2)
	want = []int{5, 3, 2}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("wanted: %v\ngot: %v\n", want, path)
	}
	// directional edges
	g = s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4})
	g.AddEdges_directional([]s.Edge{{1, 2}, {2, 3}, {3, 4}, {4, 1}})
	path = g.BredthFirst(2, 1)
	want = []int{2, 3, 4, 1}
	if !reflect.DeepEqual(path, want) {
		t.Errorf("wanted: %v\ngot: %v\n", want, path)
	}
}

func Test_Graph_GetNeighboursFor(t *testing.T) {
	g := s.Graph{}
	g.AddNodes([]int{1, 2, 3, 4, 5, 6, 99})
	es := []s.Edge{
		{1, 2}, {1, 3}, {1, 4},
		{4, 99}, {4, 5}, {4, 6},
		{99, 1}, {99, 2}, {99, 3},
	}
	g.AddEdges(es)
	got := g.GetNeighboursFor(99)
	wants := []int{1, 2, 3, 4}
	for _, w := range wants {
		if !sliceContainsInt(got, w) {
			t.Errorf("Edge to (%d) was missing", w)
		}
	}
}

func sliceContainsInt(slice []int, i int) bool {
	for _, si := range slice {
		if si == i {
			return true
		}
	}
	return false
}

func checkEdge(g s.Graph, e s.Edge) bool {
	if !sliceContainsInt(g.GetNeighboursFor(e[0]), e[1]) {
		return false
	}
	return sliceContainsInt(g.GetNeighboursFor(e[1]), e[0])
}

func checkEdge_directional(g s.Graph, e s.Edge) bool {
	return sliceContainsInt(g.GetNeighboursFor(e[0]), e[1])
}
