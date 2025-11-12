package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var puzzleData string

func main() {
	program := strings.Split(strings.TrimSpace(puzzleData), "\n")

	c := NewComputer(program...)
	c.Run()
	fmt.Printf("Part 1: %v\n", c.registers["b"])

	c.registers["a"] = 1
	c.registers["b"] = 0
	c.Run()
	fmt.Printf("Part 2: %v\n", c.registers["b"])
}

type Computer struct {
	registers          map[string]int
	memory             []string
	instructionPointer int
}

func NewComputer(program ...string) Computer {
	return Computer{
		registers: map[string]int{"a": 0, "b": 0},
		memory:    program,
	}
}

// We could certainly do more error checking, but it works fine for tests and
// real input so I'm moving on.
func (c Computer) Run() {
	for c.instructionPointer < len(c.memory) {
		curr := strings.Split(c.memory[c.instructionPointer], " ")
		switch curr[0] {
		case "hlf":
			c.registers[curr[1]] /= 2
			c.instructionPointer++
		case "tpl":
			c.registers[curr[1]] *= 3
			c.instructionPointer++
		case "inc":
			c.registers[curr[1]] += 1
			c.instructionPointer++
		case "jmp":
			offset, _ := strconv.Atoi(curr[1])
			c.instructionPointer += offset
		case "jie":
			reg := strings.TrimSuffix(curr[1], ",")
			if c.registers[reg]%2 == 0 {
				offset, _ := strconv.Atoi(curr[2])
				c.instructionPointer += offset
			} else {
				c.instructionPointer++
			}
		case "jio":
			reg := strings.TrimSuffix(curr[1], ",")
			if c.registers[reg] == 1 {
				offset, _ := strconv.Atoi(curr[2])
				c.instructionPointer += offset
			} else {
				c.instructionPointer++
			}
		}
	}
}
