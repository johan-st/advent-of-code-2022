package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/johan-st/advent-of-code-2022/util"
)

func main() {
	ins, err := instructions{}.fromString(u.Load("input.txt"))
	if err != nil {
		panic(err)
	}
	ms := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, m := range ms {
		localSum := signalAt(ins, m)
		sum += localSum
		fmt.Printf("%d: %d\n", m, localSum)
	}
	// fmt.Println(ins)
	fmt.Printf("SUM: %d\n", sum)

}

func signalAt(ins instructions, cycles int) int {
	c := cpu{ins: ins, registry: 1, cycle: 0}
	return c.run(cycles)
}

type cpu struct {
	ins      instructions
	registry int
	cycle    int
}

func (c cpu) run(measure int) int {
	for _, i := range c.ins {
		if i.addx {
			c.cycle++
			if c.cycle >= measure {
				return c.signal()
			}
			c.cycle++
			if c.cycle >= measure {
				return c.signal()
			}
			c.registry += i.value
			// fmt.Println(c.registry-i.value, "+", i.value, c.registry)
		} else {
			c.cycle++
			if c.cycle >= measure {
				return c.signal()
			}
		}
	}
	return 0
}

func (c cpu) signal() int {
	fmt.Println(c.cycle, c.registry, c.cycle*c.registry)
	return c.cycle * c.registry
}

type instruction struct {
	addx  bool
	value int
}

type instructions []instruction

func (is instructions) fromString(s string) (instructions, error) {
	rows := strings.Split(s, "\r\n")
	for _, r := range rows {
		if r[:4] == "addx" {
			v, err := strconv.Atoi(r[5:])
			if err != nil {
				return nil, err
			}
			is = append(is, instruction{addx: true, value: v})
		} else {
			is = append(is, instruction{addx: false})
		}

	}
	return is, nil
}

func (is instructions) String() string {
	str := ""
	for _, i := range is {
		if i.addx {
			str += fmt.Sprintf("addx %d\r\n", i.value)
		} else {
			str += "noop\r\n"
		}
	}
	return str
}
