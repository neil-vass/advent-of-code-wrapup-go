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
	re := regexp.MustCompile(`^To continue, please consult the code grid in the manual.  Enter the code at row (\d+), column (\d+).\n$`)
	var row, col int
	input.Parse(re, puzzleData, &row, &col)
	pos := PositionInSequence(row, col)
	fmt.Printf("Part 1: %v\n", CodeFor(pos))
}

// By inspection.
func PositionInSequence(row, col int) int {
	tri := row + col - 1
	return ((tri - 1) * tri / 2) + col
}

// I thought I'd need a clever solution here but just
// working my way up to the value I want is very fast.
func CodeFor(pos int) int {
	result := 20151125
	for i := 1; i < pos; i++ {
		result = (result * 252533) % 33554393
	}
	return result
}
