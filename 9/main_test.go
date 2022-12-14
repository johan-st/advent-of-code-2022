package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_rope_dragTail(t *testing.T) {
	tests := []struct {
		name string
		r    rope
		ra   rope
	}{
		// Same space
		{"same space 1", rope{point{0, 1}, point{0, 1}}, rope{point{0, 1}, point{0, 1}}},
		{"same space 2", rope{point{0, 0}, point{0, 0}}, rope{point{0, 0}, point{0, 0}}},
		{"same space 3", rope{point{3, 8}, point{3, 8}}, rope{point{3, 8}, point{3, 8}}},
		// Y-axis
		{"move y-axis -1", rope{point{0, 0}, point{0, 2}}, rope{point{0, 0}, point{0, 1}}},
		{"move y-axis +1", rope{point{0, 5}, point{0, 3}}, rope{point{0, 5}, point{0, 4}}},
		{"no move up y-axis", rope{point{0, 5}, point{0, 4}}, rope{point{0, 5}, point{0, 4}}},
		{"no move down y-axis", rope{point{0, 3}, point{0, 4}}, rope{point{0, 3}, point{0, 4}}},
		// X-axis
		{"move x-axis -1", rope{point{5, 0}, point{7, 0}}, rope{point{5, 0}, point{6, 0}}},
		{"move x-axis +1", rope{point{5, 0}, point{3, 0}}, rope{point{5, 0}, point{4, 0}}},
		{"no move up x-axis", rope{point{5, 0}, point{4, 0}}, rope{point{5, 0}, point{4, 0}}},
		{"no move down x-axis", rope{point{5, 0}, point{6, 0}}, rope{point{5, 0}, point{6, 0}}},
		// Diagonals
		{"diagonal no move 1", rope{point{5, 4}, point{4, 3}}, rope{point{5, 4}, point{4, 3}}},
		{"diagonal no move 2", rope{point{5, 4}, point{4, 5}}, rope{point{5, 4}, point{4, 5}}},
		{"diagonal no move 3", rope{point{5, 4}, point{6, 3}}, rope{point{5, 4}, point{6, 3}}},
		{"diagonal no move 4", rope{point{5, 4}, point{6, 5}}, rope{point{5, 4}, point{6, 5}}},
		{"diagonal up-right 1", rope{point{5, 4}, point{4, 2}}, rope{point{5, 4}, point{5, 3}}},
		{"diagonal up-right 2", rope{point{5, 4}, point{3, 3}}, rope{point{5, 4}, point{4, 4}}},
		{"diagonal down-right 1", rope{point{5, 4}, point{6, 2}}, rope{point{5, 4}, point{5, 3}}},
		{"diagonal down-right 2", rope{point{5, 4}, point{7, 3}}, rope{point{5, 4}, point{6, 4}}},
		{"diagonal up-left 1", rope{point{5, 4}, point{4, 6}}, rope{point{5, 4}, point{5, 5}}},
		{"diagonal up-left 2", rope{point{5, 4}, point{3, 5}}, rope{point{5, 4}, point{4, 4}}},
		{"diagonal down-left 1", rope{point{5, 4}, point{6, 6}}, rope{point{5, 4}, point{5, 5}}},
		{"diagonal down-left 2", rope{point{5, 4}, point{7, 5}}, rope{point{5, 4}, point{6, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.dragTail()
			if !reflect.DeepEqual(tt.r, tt.ra) {
				fmt.Printf("%s\nwanted:\n%v\ngot:\n%v", tt.name, tt.ra, tt.r)
				t.Fail()
			}
		})
	}
}

func Test_ropeSimulation_moveCase(t *testing.T) {
	sim := newSimulation(10)
	ok := testSim(sim, &rope{point{5, 5}, point{5, 5}}, 1)
	if !ok {
		t.FailNow()
	}
	sim.moveUp()
	ok = testSim(sim, &rope{point{5, 6}, point{5, 5}}, 1)
	if !ok {
		t.FailNow()
	}
	sim.moveUp()
	ok = testSim(sim, &rope{point{5, 7}, point{5, 6}}, 2)
	if !ok {
		t.FailNow()
	}
	sim.moveUp()
	ok = testSim(sim, &rope{point{5, 8}, point{5, 7}}, 3)
	if !ok {
		t.FailNow()
	}
	sim.moveDown()
	sim.moveDown()
	sim.moveDown()
	ok = testSim(sim, &rope{point{5, 5}, point{5, 6}}, 3)
	if !ok {
		t.FailNow()
	}
	sim.moveRight()
	sim.moveRight()
	sim.moveRight()
	ok = testSim(sim, &rope{point{8, 5}, point{7, 5}}, 5)
	if !ok {
		t.FailNow()
	}
}

func testSim(sim ropeSimulation, r *rope, num int) bool {
	if !reflect.DeepEqual(sim.rope, r) {
		fmt.Printf("wanted:\n%v\ngot:\n%v\n", r, sim.rope)
		return false
	}
	if sim.numVisited() != num {
		fmt.Printf("wanted:\n%v\ngot:\n%v\n", num, sim.numVisited())
		return false
	}
	return true
}
