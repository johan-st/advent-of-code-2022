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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
