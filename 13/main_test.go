package main

import (
	"reflect"
	"testing"
)

func Test_extract(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"one value", args{"9"}, []any{"9"}},
		{"multiple values", args{"1,1,3,1,1"}, []any{"1", "1", "3", "1", "1"}},
		{"nested values", args{"[1],4"}, []any{[]any{"1"}, "4"}},
		{"nested lists no values", args{"[[]]"}, []any{[]any{[]any{}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extract(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pair_correct(t *testing.T) {
	tests := []struct {
		name string
		p    pair
		want bool
	}{
		{"single value correct", pair{[]any{"1"}, []any{"2"}}, true},
		{"single value wrong", pair{[]any{"2"}, []any{"1"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.correct(); got != tt.want {
				t.Errorf("pair.correct() = %v, want %v", got, tt.want)
			}
		})
	}
}
