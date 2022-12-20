package parser_test

import (
	"reflect"
	"testing"

	p "github.com/johan-st/advent-of-code-2022/util/parser"
)

// PARSERS

func TestDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want p.Result
	}{
		{"empty string", args{""}, p.Result{"", ""}},
		{"multiple digits", args{"123456"}, p.Result{"1", "23456"}},
		{"non digit", args{"a0"}, p.Result{"", "a0"}},
		{"digit then non digit", args{"0a"}, p.Result{"0", "a"}},

		// exhaustive tests of valid runes (and "one-off"-errors)
		{"rune < digits", args{string(rune('0' - 1))}, p.Result{"", "/"}},
		{"single digit", args{"0"}, p.Result{"0", ""}},
		{"single digit", args{"1"}, p.Result{"1", ""}},
		{"single digit", args{"2"}, p.Result{"2", ""}},
		{"single digit", args{"3"}, p.Result{"3", ""}},
		{"single digit", args{"4"}, p.Result{"4", ""}},
		{"single digit", args{"5"}, p.Result{"5", ""}},
		{"single digit", args{"6"}, p.Result{"6", ""}},
		{"single digit", args{"7"}, p.Result{"7", ""}},
		{"single digit", args{"8"}, p.Result{"8", ""}},
		{"single digit", args{"9"}, p.Result{"9", ""}},
		{"rune > digits", args{string(rune('9' + 1))}, p.Result{"", ":"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Digit()(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	type args struct {
		r rune
		s string
	}
	tests := []struct {
		name string
		args args
		want p.Result
	}{
		{"empty string", args{'5', ""}, p.Result{"", ""}},
		{"single hit", args{'a', "a"}, p.Result{"a", ""}},
		{"multiple hits", args{'a', "abc"}, p.Result{"a", "bc"}},
		{"single miss", args{'.', "abc"}, p.Result{"", "abc"}},
		{"single hit then miss", args{'.', ".abc"}, p.Result{".", "abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Rune(tt.args.r)(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rune() = %v, want %v", got, tt.want)
			}
		})
	}
}

// COMBINATORS

func TestSome(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name string
		args args
		want p.Result
	}{
		{"empty string", args{p.Digit(), ""}, p.Result{"", ""}},
		{"single digit", args{p.Digit(), "1"}, p.Result{"1", ""}},
		{"multiple digits", args{p.Digit(), "123456"}, p.Result{"123456", ""}},
		{"non-digit", args{p.Digit(), "a0"}, p.Result{"", "a0"}},
		{"single digit then non-digit", args{p.Digit(), "0a"}, p.Result{"0", "a"}},
		{"multiple digits then non-digit", args{p.Digit(), "099a7"}, p.Result{"099", "a7"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Some(tt.args.p)(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOneOf(t *testing.T) {
	type args struct {
		p1 p.Parser
		p2 p.Parser
		s  string
	}
	tests := []struct {
		name string
		args args
		want p.Result
	}{
		{"empty string", args{p.Digit(), p.Rune('a'), ""}, p.Result{}},
		{"no hit", args{p.Digit(), p.Rune('a'), "bc"}, p.Result{"", "bc"}},
		{"first hit", args{p.Digit(), p.Rune('a'), "1a"}, p.Result{"1", "a"}},
		{"second hit", args{p.Digit(), p.Rune('a'), "a1"}, p.Result{"a", "1"}},
		{"both hit", args{p.Digit(), p.Rune('1'), "1a"}, p.Result{"1", "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.OneOf(tt.args.p1, tt.args.p2)(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OneOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
