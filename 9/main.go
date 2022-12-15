package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	moves, err := moves{}.fromString(u.Load("input.txt"))
	if err != nil {
		panic(err)
	}
	sim := newSimulation(5000, 10)

	sim.applyMoves(moves)
	if err != nil {
		panic(err)
	}
	// fmt.Println(sim.area) // for small tests
	fmt.Println("number of positions visited by the tail:", sim.numVisited())

}

func newSimulation(sizeArea int, lengthRope int) ropeSimulation {
	area := area{}
	for i := 0; i < sizeArea; i++ {
		area = append(area, []position{})
		for j := 0; j < sizeArea; j++ {
			area[i] = append(area[i], position{})
		}
	}
	startPos := point{sizeArea / 2, sizeArea / 2}
	r := newRope(startPos, lengthRope)
	sim := ropeSimulation{area: area, rope: r}
	sim.markTailVisit()
	return sim
}

func newRope(p point, len int) *rope {
	r := rope{}
	for i := 0; i < len; i++ {
		r = append(r, p)
	}
	return &r
}

type point struct {
	x, y int
}
type rope []point

type position struct {
	visited bool
}
type area [][]position

type move struct {
	direction rune
	steps     int
}
type moves []move
type ropeSimulation struct {
	area area
	rope *rope
}

func (sim ropeSimulation) markTailVisit() {
	r := sim.rope
	sim.area[(*r)[len(*r)-1].x][(*r)[len(*r)-1].y].visited = true
}
func (s ropeSimulation) applyMoves(ms moves) error {
	for i, m := range ms {
		err := s.applyMove(m)
		if err != nil {
			return fmt.Errorf("error on itteration %d, %e", i, err)
		}
	}
	return nil
}

func (s ropeSimulation) applyMove(m move) error {
	switch m.direction {
	case 'U':
		for i := 0; i < m.steps; i++ {
			s.moveUp()
		}
	case 'D':
		for i := 0; i < m.steps; i++ {
			s.moveDown()
		}
	case 'L':
		for i := 0; i < m.steps; i++ {
			s.moveLeft()
		}
	case 'R':
		for i := 0; i < m.steps; i++ {
			s.moveRight()
		}
	default:
		return fmt.Errorf("could not parse direction %s", string(m.direction))
	}
	return nil
}

func (s ropeSimulation) moveUp() {
	r := s.rope
	(*r)[0].y++
	s.rope.dragTail()
	s.markTailVisit()
}

func (s ropeSimulation) moveDown() {
	r := s.rope
	(*r)[0].y--
	s.rope.dragTail()
	s.markTailVisit()
}
func (s ropeSimulation) moveLeft() {
	r := s.rope
	(*r)[0].x--
	s.rope.dragTail()
	s.markTailVisit()
}
func (s ropeSimulation) moveRight() {
	r := s.rope
	(*r)[0].x++
	s.rope.dragTail()
	s.markTailVisit()
}
func (r *rope) dragTail() {
	if len(*r) < 2 {
		return
	}
	for i := 1; i < len(*r); i++ {
		updateRopePoint((*r)[i-1], &(*r)[i])
	}
}

func updateRopePoint(head point, tail *point) {
	dX := head.x - (*tail).x
	dY := head.y - (*tail).y
	if dX <= 1 && dX >= -1 && dY <= 1 && dY >= -1 {
		// tail need nor move sonce it is whithin one step of the head
		return
	}
	if dX == 0 {
		// tail is in same row as head
		if dY < 1 {
			(*tail).y--
		} else {
			(*tail).y++
		}
	}
	if dY == 0 {
		// tail is in same collumn as head
		if dX < 1 {
			(*tail).x--
		} else {
			(*tail).x++
		}

	}
	if dX != 0 && dY != 0 {
		// tail is not in the same column or row as the head and are to make a diaonal move to catch up.
		if dX > 0 && dY > 0 {
			(*tail).x++
			(*tail).y++
		} else if dX > 0 && dY < 0 {
			(*tail).x++
			(*tail).y--
		} else if dX < 0 && dY > 0 {
			(*tail).x--
			(*tail).y++
		} else {
			(*tail).x--
			(*tail).y--
		}
	}
}
func (s ropeSimulation) numVisited() int {
	sum := 0
	for _, row := range s.area {
		for _, pos := range row {
			if pos.visited {
				sum++
			}
		}
	}
	return sum
}

func (ms moves) fromString(str string) (moves, error) {
	rows := strings.Split(str, "\r\n")
	for _, r := range rows {
		dir := rune(r[0])
		if dir != 'U' && dir != 'D' && dir != 'L' && dir != 'R' {
			return ms, fmt.Errorf("move could not parse direction %s", string(dir))
		}

		numStr := r[2:]

		steps, err := strconv.Atoi(numStr)
		if err != nil {
			return ms, fmt.Errorf("move could not parse number %s", numStr)
		}
		ms = append(ms, move{direction: rune(dir), steps: steps})
	}
	return ms, nil
}

func (a area) String() string {
	str := ""
	for _, rows := range a {
		for _, pos := range rows {
			if pos.visited {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}
