package main

import (
	"reflect"
	"testing"
)

func Test_scoreDuplicate(t *testing.T) {
	type args struct {
		d rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a = 1", args{'a'}, 1},
		{"b = 2", args{'b'}, 2},
		{"c = 3", args{'c'}, 3},
		{"d = 4", args{'d'}, 4},
		{"e = 5", args{'e'}, 5},
		{"f = 6", args{'f'}, 6},
		{"g = 7", args{'g'}, 7},
		{"h = 8", args{'h'}, 8},
		{"i = 9", args{'i'}, 9},
		{"j = 10", args{'j'}, 10},
		{"k = 11", args{'k'}, 11},
		{"l = 12", args{'l'}, 12},
		{"m = 13", args{'m'}, 13},
		{"n = 14", args{'n'}, 14},
		{"o = 15", args{'o'}, 15},
		{"p = 16", args{'p'}, 16},
		{"q = 17", args{'q'}, 17},
		{"r = 18", args{'r'}, 18},
		{"s = 19", args{'s'}, 19},
		{"t = 20", args{'t'}, 20},
		{"u = 21", args{'u'}, 21},
		{"v = 22", args{'v'}, 22},
		{"w = 23", args{'w'}, 23},
		{"x = 24", args{'x'}, 24},
		{"y = 25", args{'y'}, 25},
		{"z = 26", args{'z'}, 26},
		{"A = 27", args{'A'}, 27},
		{"B = 28", args{'B'}, 28},
		{"C = 29", args{'C'}, 29},
		{"D = 30", args{'D'}, 30},
		{"E = 31", args{'E'}, 31},
		{"F = 32", args{'F'}, 32},
		{"G = 33", args{'G'}, 33},
		{"H = 34", args{'H'}, 34},
		{"I = 35", args{'I'}, 35},
		{"J = 36", args{'J'}, 36},
		{"K = 37", args{'K'}, 37},
		{"L = 38", args{'L'}, 38},
		{"M = 39", args{'M'}, 39},
		{"N = 40", args{'N'}, 40},
		{"O = 41", args{'O'}, 41},
		{"P = 42", args{'P'}, 42},
		{"Q = 43", args{'Q'}, 43},
		{"R = 44", args{'R'}, 44},
		{"S = 45", args{'S'}, 45},
		{"T = 46", args{'T'}, 46},
		{"U = 47", args{'U'}, 47},
		{"V = 48", args{'V'}, 48},
		{"W = 49", args{'W'}, 49},
		{"X = 50", args{'X'}, 50},
		{"Y = 51", args{'Y'}, 51},
		{"Z = 52", args{'Z'}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreDuplicate(tt.args.d); got != tt.want {
				t.Errorf("scoreDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicates(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"empty string", args{[]string{""}}, []rune{}},
		{"single character", args{[]string{"a"}}, []rune{}},
		{"two characters", args{[]string{"a", "b"}}, []rune{}},
		{"two characters", args{[]string{"a", "a"}}, []rune{'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
