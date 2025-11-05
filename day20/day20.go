package main

import (
	_ "embed"
	"fmt"
	"regexp"

	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

//go:embed input.txt
var puzzleData string

func main() {
	var target int
	input.Parse(regexp.MustCompile(`^(\d+)\n$`), puzzleData, &target)
	fmt.Printf("Part 1: %d\n", target)
}
