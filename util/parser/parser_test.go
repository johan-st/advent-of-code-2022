// Rudimentary Parser Combinator package

package parser

import (
	"reflect"
	"testing"
)

func TestParseDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []result
	}{
		{"empty string", args{""}, []result{}},
		{"multiple digits", args{"123456"}, []result{{"1", "23456"}}},
		{"non digit", args{"a0"}, []result{{"", "a0"}}},
		{"digit then non digit", args{"0a"}, []result{{"0", "a"}}},
		// exhaustive test of valid runes (and "one-off"-errors)
		{"rune < digits", args{string(rune('0' - 1))}, []result{{"", "/"}}},
		{"single digit", args{"0"}, []result{{"0", ""}}},
		{"single digit", args{"1"}, []result{{"1", ""}}},
		{"single digit", args{"2"}, []result{{"2", ""}}},
		{"single digit", args{"3"}, []result{{"3", ""}}},
		{"single digit", args{"4"}, []result{{"4", ""}}},
		{"single digit", args{"5"}, []result{{"5", ""}}},
		{"single digit", args{"6"}, []result{{"6", ""}}},
		{"single digit", args{"7"}, []result{{"7", ""}}},
		{"single digit", args{"8"}, []result{{"8", ""}}},
		{"single digit", args{"9"}, []result{{"9", ""}}},
		{"rune > digits", args{string(rune('9' + 1))}, []result{{"", ":"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDigit(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	type args struct {
		p parser
		s string
	}
	tests := []struct {
		name string
		args args
		want []result
	}{
		{"empty string", args{ParseDigit, ""}, []result{}},
		{"single digit", args{ParseDigit, "1"}, []result{{"1", ""}}},
		{"multiple digits", args{ParseDigit, "123456"}, []result{{"123456", ""}}},
		{"non-digit", args{ParseDigit, "a0"}, []result{{"", "a0"}}},
		{"single digit then non-digit", args{ParseDigit, "0a"}, []result{{"0", "a"}}},
		{"multiple digits then non-digit", args{ParseDigit, "099a7"}, []result{{"099", "a7"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.p, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}
