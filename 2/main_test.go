package main

import (
	"testing"
)

func Test_scoreGame(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{""}, 0},
		{"A X", args{"A X"}, 4},
		{"A Y", args{"A Y"}, 8},
		{"A Z", args{"A Z"}, 3},
		{"B X", args{"B X"}, 1},
		{"B Y", args{"B Y"}, 5},
		{"B Z", args{"B Z"}, 9},
		{"C X", args{"C X"}, 7},
		{"C Y", args{"C Y"}, 2},
		{"C Z", args{"C Z"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreGame(tt.args.in); got != tt.want {
				t.Errorf("scoreGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreGamePart2(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{""}, 0},
		{"A X", args{"A X"}, 3},
		{"A Y", args{"A Y"}, 4},
		{"A Z", args{"A Z"}, 8},
		{"B X", args{"B X"}, 1},
		{"B Y", args{"B Y"}, 5},
		{"B Z", args{"B Z"}, 9},
		{"C X", args{"C X"}, 2},
		{"C Y", args{"C Y"}, 6},
		{"C Z", args{"C Z"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreGamePart2(tt.args.game); got != tt.want {
				t.Errorf("scoreGamePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
