package main

// u "github.com/johan-st/advent-of-code-2022/util"
// search "github.com/johan-st/advent-of-code-2022/util/search"
// p "github.com/johan-st/advent-of-code-2022/util/parser"

import (
	"fmt"
	"regexp"
	"strings"

	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	sig := signal{}
	sig.fromString(u.Load("sample.txt"))

	fmt.Printf("%v\n", sig)
	for i, p := range sig.pairs {
		fmt.Printf("%d: %v\n", i+1, p.correct())
	}

}

type signal struct {
	pairs []pair
}
type pair struct {
	left, right []any
}

func (s *signal) fromString(in string) {
	splitIn := strings.Split(in, "\r\n\r\n")
	for _, pairStr := range splitIn {
		p := pair{}
		p.fromString(pairStr)
		s.pairs = append(s.pairs, p)
	}
}
func (s signal) String() string {
	str := ""
	for _, p := range s.pairs {
		str += p.String() + "\n"
	}
	return str
}

func (p pair) correct() bool {
	compRes := compare(p.left, p.right)
	switch compRes {
	case "left":
		return true
	default:
		panic("unexpected response from compare")

	}
}

func compare(left, right any) string {
	undecided := false
	itt := 0
	for undecided {
		itt++
		if itt > 10 {
			panic("something went wrong in the compare function")
		}

	}
	return "left"
}

func (p *pair) fromString(in string) {
	splitIn := strings.Split(in, "\r\n")
	p.left = extract(splitIn[0][1 : len(splitIn[0])-1])
	p.right = extract(splitIn[1][1 : len(splitIn[1])-1])
}

func (p pair) String() string {
	return fmt.Sprintf("%v\n%v\n", p.left, p.right)
}

func extract(str string) []any {
	// fmt.Printf("extracting from string %s\n", str)
	RElist := regexp.MustCompile(`\[(.*)\]`)
	REnums := regexp.MustCompile(`\d`)

	listIndecies := RElist.FindStringSubmatchIndex(str)
	// fmt.Printf("%d\n", len(listIndecies))
	if len(listIndecies) == 4 {
		list := str[listIndecies[2]:listIndecies[3]]
		before := REnums.FindAllString(str[:listIndecies[2]], 100)
		after := REnums.FindAllString(str[listIndecies[3]:], 100)
		ret := []any{}
		for _, val := range before {
			ret = append(ret, val)
		}
		ret = append(ret, extract(list))
		for _, val := range after {
			ret = append(ret, val)
		}
		// fmt.Printf("%v\n", ret)

		return ret
	}

	nums := REnums.FindAllString(str, 100)
	ret := []any{}
	for _, num := range nums {
		ret = append(ret, num)
	}
	// fmt.Printf("%v\n", ret)
	return ret

}
