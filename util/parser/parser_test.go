package parser_test

import (
	"reflect"
	"testing"

	p "github.com/johan-st/advent-of-code-2022/util/parser"
)

// PARSERS

// DEBUG CODE GOES HERE
// TODO: REMOVE BEFORE COMMIT

// DEBUG END

func TestDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{""}, p.Result{}, true},
		{"rune before 0", args{"/"}, p.Result{}, true},
		{"0", args{"0"}, []any{"0"}, false},
		{"1", args{"1"}, p.Result{"1"}, false},
		{"2", args{"2"}, p.Result{"2"}, false},
		{"3", args{"3"}, p.Result{"3"}, false},
		{"4", args{"4"}, p.Result{"4"}, false},
		{"5", args{"5"}, p.Result{"5"}, false},
		{"6", args{"6"}, p.Result{"6"}, false},
		{"7", args{"7"}, p.Result{"7"}, false},
		{"8", args{"8"}, p.Result{"8"}, false},
		{"9", args{"9"}, p.Result{"9"}, false},
		{"rune after 9", args{":"}, p.Result{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(p.Digit(), tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Digit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digit() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestDigits(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.Digits(), ""}, p.Result{}, true},
		{"match", args{p.Digits(), "12a"}, p.Result{"1", "2"}, false},
		{"miss", args{p.Digits(), "b12"}, p.Result{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Digits() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.Digit(), ""}, p.Result{}, false},
		{"single digit", args{p.Digit(), "1"}, p.Result{"1"}, false},
		{"multiple digits", args{p.Digit(), "123"}, p.Result{"1", "2", "3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(p.Some(tt.args.p), tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Some() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.Rune('a'), ""}, p.Result{}, true},
		{"match", args{p.Rune('a'), "ab"}, p.Result{"a"}, false},
		{"miss", args{p.Rune('a'), "ba"}, p.Result{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rune() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.Int(), ""}, p.Result{}, true},
		{"miss, malformed int", args{p.Int(), "--42"}, p.Result{}, true},
		{"miss, letter in first pos", args{p.Int(), "a42"}, p.Result{}, true},
		{"hit positive", args{p.Int(), "42"}, p.Result{42}, false},
		{"hit negative", args{p.Int(), "-42"}, p.Result{-42}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOneOf(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.OneOf([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), ""}, p.Result{}, true},
		{"first is a match", args{p.OneOf([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "abc"}, p.Result{"a"}, false},
		{"second is a match", args{p.OneOf([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "bca"}, p.Result{"b"}, false},
		{"no match", args{p.OneOf([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "d"}, p.Result{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("OneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OneOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequence(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), ""}, p.Result{}, true},
		{"match", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "abcd"}, p.Result{"a", "b", "c"}, false},
		{"missing first", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "bcd"}, p.Result{}, true},
		{"missing last", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "ab"}, p.Result{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sequence() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAll(t *testing.T) {
	type args struct {
		p p.Parser
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    p.Result
		wantErr bool
	}{
		{"empty string", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), ""}, p.Result{}, true},
		{"error, input left", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "abcd"}, p.Result{}, true},
		{"missing first", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "bcd"}, p.Result{}, true},
		{"missing last", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "ab"}, p.Result{}, true},
		{"exact match", args{p.Sequence([]p.Parser{p.Rune('a'), p.Rune('b'), p.Rune('c')}), "abc"}, p.Result{"a", "b", "c"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.ParseAll(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAll() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
