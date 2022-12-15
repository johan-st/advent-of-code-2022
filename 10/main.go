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

	crt := crtScreen{}
	c := cpu{ins: ins, registry: 1}
	for i := 0; i < 240; i++ {
		crt.update(c)
		c.tick()
		// fmt.Println(c)
	}
	fmt.Println(crt)
}

type crtScreen struct {
	rows [6][40]rune
}

func (crt *crtScreen) update(c cpu) {
	line := c.cycle / 40
	char := c.cycle % 40

	if c.registry >= char-1 && c.registry <= char+1 {
		crt.rows[line][char] = 'X'
	} else {
		crt.rows[line][char] = '.'
	}

}

func (crt crtScreen) String() string {
	str := "------------------------------------------\n"
	for _, row := range crt.rows {
		str += "|"
		for _, ru := range row {
			if string(ru) != "\x00" {
				str += string(ru)
			} else {
				str += "."
			}
		}
		str += "|\n"
	}
	str += "------------------------------------------"
	return str
}

type cpu struct {
	ins                instructions
	registry           int
	cycle              int
	cycleOnInstruction int
	currentInstruction int
}

func (c *cpu) tick() {
	c.cycle++
	if c.ins[c.currentInstruction].addx && c.cycleOnInstruction == 0 {
		c.cycleOnInstruction++
	} else if c.ins[c.currentInstruction].addx && c.cycleOnInstruction == 1 {
		c.registry += c.ins[c.currentInstruction].value
		c.currentInstruction++
		c.cycleOnInstruction = 0
	} else {
		c.currentInstruction++
	}
}

func (c cpu) signal() int {
	fmt.Println(c.cycle, c.registry, c.cycle*c.registry)
	return c.cycle * c.registry
}
func (c cpu) String() string {
	return fmt.Sprintf("--CPU--\ncycle: %d\nins: %d of %d\nreg:%d\nregCycle:%d\n-------\n", c.cycle, len(c.ins), c.currentInstruction, c.registry, c.cycleOnInstruction)
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
