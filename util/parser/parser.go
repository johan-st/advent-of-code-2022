// Rudimentary Parser Combinator package
package parser

import "fmt"

/*
A Parser for things,
Is a function from string,
To lists of pairs,
Of things and strings.
*/
type Parser func(string) Result

type Result struct {
	Parsed    string
	Remainder string
}

// PARSERS

// parse a single digit
func Digit() Parser {
	return func(s string) Result {
		if len(s) < 1 {
			return Result{"", s}
		}
		if s[0] >= '0' && s[0] <= '9' {
			return Result{string(s[0]), s[1:]}
		}
		return Result{"", s}
	}
}

// parse a single rune
func Rune(r rune) Parser {
	return func(s string) Result {
		if len(s) < 1 {
			return Result{"", s}
		}
		if rune(s[0]) == r {
			return Result{string(s[0]), s[1:]}
		}
		return Result{"", s}
	}
}

// Parse an integer
func Int() Parser {
	return Pipe([]Parser{
		Rune('-'),
		Some(Digit()),
	})
}

// COMBINATORS

// parse zero or more
func Some(p Parser) Parser {
	return func(s string) Result {
		// first itteration
		parsed := p(s)
		retResult := parsed.Parsed
		s = parsed.Remainder

		// remaining itterations
		for parsed.Parsed != "" {
			parsed = p(s)
			if len(parsed.Parsed) < 1 {
				continue
			}
			retResult += parsed.Parsed
			s = parsed.Remainder
		}
		return Result{retResult, s}
	}
}

// match either the first or the second parser
func OneOf(p1 Parser, p2 Parser) Parser {
	return func(s string) Result {
		if len(s) < 1 {
			return Result{}
		}
		fmt.Println("string is", s)
		if res := p1(s); len(res.Parsed) > 0 {
			fmt.Println("FIRST\n", res)
			return res
		}
		if res := p2(s); len(res.Parsed) > 0 {
			fmt.Println("SECOND\n", res)
			return res
		}
		return Result{"", s}
	}
}

// Run the string through all parsers in order
func Pipe(ps []Parser) Parser {
	return func(s string) Result {
		res := Result{"", s}
		for _, p := range ps {
			r := p(res.Remainder)
			res = Result{res.Parsed + r.Parsed, r.Remainder}
		}
		return res
	}
}
