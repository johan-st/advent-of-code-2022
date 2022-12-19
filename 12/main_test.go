package main

import (
	"fmt"
	"reflect"
	"testing"

	s "github.com/johan-st/advent-of-code-2022/util/search"
)

func Test_graphFromMap(t *testing.T) {
	heightMap := mapFromString(sampleString())
	g, s, e := graphFromMap(heightMap)
	if s != 1 {
		fmt.Println("starting position is wrong")
		t.Fail()
	}
	if e != 22 {
		fmt.Println("end position is wrong")
		t.Fail()
	}

	if reflect.DeepEqual(g, sampleGraph()) {
		fmt.Printf("%v\n", sampleGraph())
		t.Fail()
	}
}

func sampleString() string {
	return "Sabqponm\r\nabcryxxl\r\naccszExk\r\nacctuvwj\r\nabdefghi"
}
func sampleGraph() s.Graph {
	g := s.Graph{}
	ids := []int{}
	for id := 1; id <= 40; id++ {
		ids = append(ids, id)
	}
	g.AddNodes(ids)
	g.AddEdges_directional(
		[]s.Edge{
			{1, 9}, {1, 2}, {2, 10},
			{2, 1}, {2, 3},
			{3, 11}, {3, 2},
			{4, 12}, {4, 3}, {4, 5},
			{5, 4}, {5, 6},
			{6, 5}, {6, 7},
			{7, 6},
			{8, 16}, {8, 7},
			{9, 1}, {9, 17}, {9, 10},
			{10, 2}, {10, 18}, {10, 9}, {10, 11},
			{11, 3}, {11, 19}, {11, 10},
			{12, 4}, {12, 20}, {12, 11},
			{13, 5}, {13, 21}, {13, 12}, {13, 14},
			{14, 6}, {14, 13}, {14, 15},
			{15, 7}, {15, 23}, {15, 14},
			{16, 8}, {16, 24},
			{17, 9}, {17, 25},
			{18, 10}, {18, 26}, {18, 17}, {18, 19},
			{19, 11}, {19, 27}, {19, 18},
			{20, 12}, {20, 28}, {20, 19},
			{21, 13}, {21, 29}, {21, 20}, {21, 22},
			{22, 14}, {22, 30}, {22, 21}, {22, 23},
			{23, 15}, {23, 31},
			{24, 16}, {24, 32},
			{25, 17},
			{26, 18}, {26, 25}, {26, 27},
			{27, 19}, {27, 26},
			{28, 20}, {28, 27}, {28, 29},
			{29, 28}, {29, 30},
			{30, 29}, {30, 31},
			{31, 23}, {31, 30},
			{32, 24},
			{33, 25}, {33, 34},
			{34, 26}, {34, 33},
			{35, 27}, {35, 34}, {35, 36},
			{36, 35}, {36, 37},
			{37, 36}, {37, 38},
			{38, 37}, {38, 39},
			{39, 38},
			{40, 32}, {40, 39}},
	)
	return g
}

func Test_validNeighbours(t *testing.T) {
	type args struct {
		h heightMap
		p pos
	}
	tests := []struct {
		name string
		args args
		want []pos
	}{
		{"first position", args{mapFromString(sampleString()), pos{0, 0}}, []pos{{1, 0}, {0, 1}}},
		{"last position", args{mapFromString(sampleString()), pos{4, 7}}, []pos{{3, 7}, {4, 6}}},
		{"39 position", args{mapFromString(sampleString()), pos{4, 6}}, []pos{{4, 5}, {4, 7}}},
		{"top/end", args{mapFromString(sampleString()), pos{2, 4}}, []pos{{1, 4}, {3, 4}, {2, 3}, {2, 5}}},
		// {"new", args{mapFromString(sampleString()), pos{}}, []pos{{}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validNeighbours(tt.args.h, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validNeighbours(hm, %v) = %v, want %v", tt.args.p, got, tt.want)
			}
		})
	}
}
