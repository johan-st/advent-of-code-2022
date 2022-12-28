// Rudimentary Parser Combinator package
package parser

import (
	"fmt"
	"strconv"
)

/*
A Parser for things,
Is a function from string,
To lists of pairs,
Of things and strings.
*/
type Parser func(State) (State, error)

type State struct {
	index int
	input string
	Res   Result
}
type Result []any

// BASE PARSERS

// parse a single digit
func Digit() Parser {
	return func(s State) (State, error) {
		if len(s.input) <= s.index {
			return s, fmt.Errorf("expected a Digit at index %d but got unexpected end of input", s.index)
		}
		b := s.input[s.index]
		// numbers are in this byte-range
		if b < 0x30 || b > 0x39 {
			return s, fmt.Errorf("expected a Digit at index %d but got %s", s.index, string(b))
		}
		return updateState(s, 1, string(b)), nil
	}
}

// parse a single rune
func Rune(r rune) Parser {
	return func(s State) (State, error) {
		if len(s.input) <= s.index {
			return s, fmt.Errorf("expected the rune '%s' at index %d but got unexpected end of input", string(r), s.index)
		}
		b := s.input[s.index]
		if rune(b) != r {
			return s, fmt.Errorf("expected the rune '%s' at index %d but got %s", string(r), s.index, string(b))
		}
		return updateState(s, 1, string(b)), nil
	}
}

// EXTENDED PARSERS

// one or more digits
func Digits() Parser {
	return Some1(Digit())
}

// Parse an integer
func Int() Parser {
	return OneOf([]Parser{
		Sequence([]Parser{
			Rune('-'),
			Digits(),
		}),
		Digits(),
	}).Concat().Map(toInt)
}

// COMBINATORS

// parse zero or more. Never fails
func Some(p Parser) Parser {
	return func(s State) (State, error) {
		err := error(nil)
		for err == nil {
			s, err = p(s)
		}
		return s, nil
	}
}

// parse one or more
func Some1(p Parser) Parser {
	return func(s State) (State, error) {
		originalState := s
		err := error(nil)
		gotResult := false
		for err == nil {
			s, err = p(s)
			if err == nil && !gotResult {
				gotResult = true
			}
		}
		if !gotResult {
			return originalState, fmt.Errorf("parser Some1 could not match anything¨at index %d", s.index)
		}
		return s, nil
	}
}

// Try several parsers and return the first match
func OneOf(ps []Parser) Parser {
	return func(s State) (State, error) {
		for _, p := range ps {
			if ns, err := p(s); err == nil {
				return ns, nil
			}
		}
		return s, fmt.Errorf("no parser out of the %d provided could parse index %d", len(ps), s.index)
	}

}

// Run the string through all parsers in order. Returns an error if any of them failes
func Sequence(ps []Parser) Parser {
	return func(s State) (State, error) {
		s2 := s
		for _, p := range ps {
			var err error
			s2, err = p(s2)
			if err != nil {
				return s, fmt.Errorf("sequence failed to parse with error: %s", err)
			}
		}
		return s2, nil
	}
}

// UTILITIES

// Parse Parser
func Parse(p Parser, str string) (Result, error) {
	state := State{index: 0, input: str, Res: Result{}}
	res, err := p(state)
	return res.Res, err
}

// Run a function on all instances of result
func (p Parser) Map(fn func(any) any) Parser {
	return func(s State) (State, error) {
		s2, err := p(s)
		if err != nil {
			return s, err
		}
		newRes := Result{}
		for _, r := range s2.Res {
			newRes = append(newRes, fn(r))
		}

		s2.Res = newRes
		return s2, nil
	}
}

func (p Parser) Concat() Parser {
	return func(s State) (State, error) {
		newS, err := p(s)
		if err != nil {
			return s, err
		}
		var conStr string
		for i, v := range newS.Res {
			switch v := v.(type) {
			case string:
				conStr += v
			default:
				return s, fmt.Errorf("tried to concatenate but encounter value of type %T at Result Index %d", v, i)
			}
		}
		newS.Res = []any{conStr}
		return newS, nil
	}
}

// get a updated state
func updateState(s State, indexChange int, value any) State {
	return State{s.index + indexChange, s.input, append(s.Res, value)}
}

func toInt(s any) any {
	switch s := s.(type) {
	case string:
		i, err := strconv.Atoi(s)
		if err != nil {
			str := fmt.Sprintf("Error in Int parser. error was: %s", err)
			panic(str)
		}
		return i
	default:
		return s
	}
}
