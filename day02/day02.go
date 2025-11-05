package main

import (
	_ "embed"
	"fmt"
	"iter"
	"regexp"
	"slices"

	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

//go:embed input.txt
var puzleInput string

type Present struct{ length, width, height int }

func PaperNeeded(p Present) int {
	sides := []int{p.length * p.width, p.width * p.height, p.height * p.length}
	paper := slices.Min(sides)
	for _, s := range sides {
		paper += 2 * s
	}
	return paper
}

func RibbonNeeded(p Present) int {
	perimeters := []int{2 * (p.length + p.width), 2 * (p.width + p.height), 2 * (p.height + p.length)}
	ribbon := slices.Min(perimeters)
	ribbon += p.length * p.width * p.height
	return ribbon
}

func ParsePresent(s string) Present {
	re := regexp.MustCompile(`^(\d+)x(\d+)x(\d+)$`)
	var p Present
	err := input.Parse(re, s, &p.length, &p.width, &p.height)
	if err != nil {
		panic(err)
	}
	return p
}

func Solve(lines iter.Seq[string], calc func(Present) int) int {
	var total int
	for ln := range lines {
		total += calc(ParsePresent(ln))
	}
	return total
}

func main() {
	fmt.Printf("Part 1: %d\n", Solve(input.Lines(puzleInput), PaperNeeded))
	fmt.Printf("Part 2: %d\n", Solve(input.Lines(puzleInput), RibbonNeeded))

}
