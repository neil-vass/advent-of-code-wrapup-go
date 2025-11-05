package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func FindFloor(s string) (floor int) {
	return strings.Count(s, "(") - strings.Count(s, ")")
}

func PosForBasement(s string) int {
	floor := 0
	for i, ch := range s {
		switch ch {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	panic("Never reached basement!")
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Printf("Part 1: %v\n", FindFloor(input))
	fmt.Printf("Part 2: %v\n", PosForBasement(input))
}
