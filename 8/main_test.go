package main

import (
	"reflect"
	"testing"
)

func newT(height int) tree {
	return tree{height: height}
}
func sampleString() string {
	return "30373\r\n25512\r\n65332\r\n33549\r\n35390"
}
func directionalForests() [4]forest {
	sampleForestLeft := forest{[]tree{newT(2), newT(2), newT(2)}, []tree{newT(1), newT(2), newT(2)}, []tree{newT(2), newT(2), newT(2)}}
	sampleForestRight := forest{[]tree{newT(2), newT(2), newT(2)}, []tree{newT(2), newT(2), newT(1)}, []tree{newT(2), newT(2), newT(2)}}
	sampleForestUp := forest{[]tree{newT(2), newT(1), newT(2)}, []tree{newT(2), newT(2), newT(2)}, []tree{newT(2), newT(2), newT(2)}}
	sampleForestDown := forest{[]tree{newT(2), newT(2), newT(2)}, []tree{newT(2), newT(2), newT(2)}, []tree{newT(2), newT(1), newT(2)}}
	return [4]forest{sampleForestLeft, sampleForestRight, sampleForestUp, sampleForestDown}
}

func sampleForest() forest {
	return forest{[]tree{newT(3), newT(0), newT(3), newT(7), newT(3)}, []tree{newT(2), newT(5), newT(5), newT(1), newT(2)}, []tree{newT(6), newT(5), newT(3), newT(3), newT(2)}, []tree{newT(3), newT(3), newT(5), newT(4), newT(9)}, []tree{newT(3), newT(5), newT(3), newT(9), newT(0)}}
}
func Test_forest_String(t *testing.T) {
	f := sampleForest()
	want := sampleString()
	got := f.String()
	if got != want {
		t.Errorf("forest.String() = %v, want %v", got, want)
	}
}

func Test_forest_fromString(t *testing.T) {
	s := sampleString()
	want := sampleForest()
	got := forest{}.fromString(s)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("forest.fromString() = %v, want %v", got, want)
	}

}

func Test_forest_isVisible(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want bool
	}{
		{"0,0", sampleForest(), args{0, 0}, true},
		{"0,1", sampleForest(), args{0, 1}, true},
		{"0,2", sampleForest(), args{0, 2}, true},
		{"0,3", sampleForest(), args{0, 3}, true},
		{"0,4", sampleForest(), args{0, 4}, true},
		{"1,0", sampleForest(), args{1, 0}, true},
		{"1,1", sampleForest(), args{1, 1}, true},
		{"1,2", sampleForest(), args{1, 2}, true},
		{"1,3", sampleForest(), args{1, 3}, false},
		{"1,4", sampleForest(), args{1, 4}, true},
		{"2,0", sampleForest(), args{2, 0}, true},
		{"2,1", sampleForest(), args{2, 1}, true},
		{"2,2", sampleForest(), args{2, 2}, false},
		{"2,3", sampleForest(), args{2, 3}, true},
		{"2,4", sampleForest(), args{2, 4}, true},
		{"3,0", sampleForest(), args{3, 0}, true},
		{"3,1", sampleForest(), args{3, 1}, false},
		{"3,2", sampleForest(), args{3, 2}, true},
		{"3,3", sampleForest(), args{3, 3}, false},
		{"3,4", sampleForest(), args{3, 4}, true},
		{"4,0", sampleForest(), args{4, 0}, true},
		{"4,1", sampleForest(), args{4, 1}, true},
		{"4,2", sampleForest(), args{4, 2}, true},
		{"4,3", sampleForest(), args{4, 3}, true},
		{"4,4", sampleForest(), args{4, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.isVisible(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.isVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forest_isVisibleLeft(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want bool
	}{

		{"visible left", directionalForests()[0], args{1, 1}, true},
		{"visible right", directionalForests()[1], args{1, 1}, false},
		{"visible up", directionalForests()[2], args{1, 1}, false},
		{"visible down", directionalForests()[3], args{1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.isVisibleLeft(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.isVisibleLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forest_isVisibleRight(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want bool
	}{

		{"visible left", directionalForests()[0], args{1, 1}, false},
		{"visible right", directionalForests()[1], args{1, 1}, true},
		{"visible up", directionalForests()[2], args{1, 1}, false},
		{"visible down", directionalForests()[3], args{1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.isVisibleRight(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.isVisibleRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forest_isVisibleUp(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want bool
	}{

		{"visible left", directionalForests()[0], args{1, 1}, false},
		{"visible right", directionalForests()[1], args{1, 1}, false},
		{"visible up", directionalForests()[2], args{1, 1}, true},
		{"visible down", directionalForests()[3], args{1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.isVisibleUp(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.isVisibleUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forest_isVisibleDown(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want bool
	}{

		{"visible left", directionalForests()[0], args{1, 1}, false},
		{"visible right", directionalForests()[1], args{1, 1}, false},
		{"visible up", directionalForests()[2], args{1, 1}, false},
		{"visible down", directionalForests()[3], args{1, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.isVisibleDown(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.isVisibleDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forest_scenicValueAt(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want int
	}{
		{"0,0", sampleForest(), args{0, 0}, 0},
		{"0,1", sampleForest(), args{0, 1}, 0},
		{"0,2", sampleForest(), args{0, 2}, 0},
		{"0,3", sampleForest(), args{0, 3}, 0},
		{"0,4", sampleForest(), args{0, 4}, 0},
		{"1,0", sampleForest(), args{1, 0}, 0},
		{"1,1", sampleForest(), args{1, 1}, 1},
		{"1,2", sampleForest(), args{1, 2}, 4},
		{"1,3", sampleForest(), args{1, 3}, 1},
		{"1,4", sampleForest(), args{1, 4}, 0},
		{"2,0", sampleForest(), args{2, 0}, 0},
		{"2,1", sampleForest(), args{2, 1}, 6},
		{"2,2", sampleForest(), args{2, 2}, 1},
		{"2,3", sampleForest(), args{2, 3}, 2},
		{"2,4", sampleForest(), args{2, 4}, 0},
		{"3,0", sampleForest(), args{3, 0}, 0},
		{"3,1", sampleForest(), args{3, 1}, 1},
		{"3,2", sampleForest(), args{3, 2}, 8},
		{"3,3", sampleForest(), args{3, 3}, 3},
		{"3,4", sampleForest(), args{3, 4}, 0},
		{"4,0", sampleForest(), args{4, 0}, 0},
		{"4,1", sampleForest(), args{4, 1}, 0},
		{"4,2", sampleForest(), args{4, 2}, 0},
		{"4,3", sampleForest(), args{4, 3}, 0},
		{"4,4", sampleForest(), args{4, 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.scenicValueAt(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.scenicValueAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forest_scenicLeft(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		f    forest
		args args
		want int
	}{
		{"0,0", sampleForest(), args{0, 0}, 0},
		{"0,1", sampleForest(), args{0, 1}, 1},
		{"0,2", sampleForest(), args{0, 2}, 2},
		{"0,3", sampleForest(), args{0, 3}, 3},
		{"0,4", sampleForest(), args{0, 4}, 1},
		{"1,0", sampleForest(), args{1, 0}, 0},
		{"1,1", sampleForest(), args{1, 1}, 1},
		{"1,2", sampleForest(), args{1, 2}, 1},
		{"1,3", sampleForest(), args{1, 3}, 1},
		{"1,4", sampleForest(), args{1, 4}, 2},
		{"2,0", sampleForest(), args{2, 0}, 0},
		{"2,1", sampleForest(), args{2, 1}, 1},
		{"2,2", sampleForest(), args{2, 2}, 1},
		{"2,3", sampleForest(), args{2, 3}, 1},
		{"2,4", sampleForest(), args{2, 4}, 1},
		{"3,0", sampleForest(), args{3, 0}, 0},
		{"3,1", sampleForest(), args{3, 1}, 1},
		{"3,2", sampleForest(), args{3, 2}, 2},
		{"3,3", sampleForest(), args{3, 3}, 1},
		{"3,4", sampleForest(), args{3, 4}, 4},
		{"4,0", sampleForest(), args{4, 0}, 0},
		{"4,1", sampleForest(), args{4, 1}, 1},
		{"4,2", sampleForest(), args{4, 2}, 1},
		{"4,3", sampleForest(), args{4, 3}, 3},
		{"4,4", sampleForest(), args{4, 4}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.scenicLeft(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("forest.scenicLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
