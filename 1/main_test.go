package main

import (
	"reflect"
	"testing"
)

func Test_splitInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty", args{""}, nil},
		{"one elf", args{"500\n500\n100"}, []string{"500\n500\n100"}},
		{"two elf", args{"500\n500\n100\n\n500\n100"}, []string{"500\n500\n100", "500\n100"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumString(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty case", args{""}, 0},
		{"one elf", args{"500\n500\n100"}, 1100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumString(tt.args.in0); got != tt.want {
				t.Errorf("sumString() = %v, want %v", got, tt.want)
			}
		})
	}
}
