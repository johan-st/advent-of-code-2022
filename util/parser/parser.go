// Rudimentary Parser Combinator package
package parser

/*
A parser for things,
Is a function from string,
To lists of pairs,
Of things and strings.
*/
type parser func(string) []result

type result struct {
	result    string
	remainder string
}

// parse a single digit
func ParseDigit(s string) []result {
	if len(s) < 1 {
		return []result{}
	}
	if s[0] >= '0' && s[0] <= '9' {
		return []result{{string(s[0]), s[1:]}}
	}
	return []result{{"", s}}
}

// parse 0 or more
func Some(p parser, s string) []result {
	// first itteration
	parsed := p(s)
	if len(parsed) < 1 {
		return parsed
	}
	retResult := parsed[0].result
	s = parsed[0].remainder

	// remaining itterations
	for len(parsed) > 0 && parsed[0].result != "" {
		parsed = p(s)
		if len(parsed) < 1 {
			continue
		}
		retResult += parsed[0].result
		s = parsed[0].remainder
	}

	return []result{{retResult, s}}
}
