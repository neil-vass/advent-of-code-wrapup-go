package main

import (
	_ "embed"
	"fmt"
)

type Pos struct{ x, y int }

func Deliveries(s string, actors int) int {
	positions := make([]Pos, actors)
	visited := map[Pos]bool{{0, 0}: true}
	var i int
	for _, move := range s {
		switch move {
		case '^':
			positions[i].y += 1
		case 'v':
			positions[i].y -= 1
		case '<':
			positions[i].x -= 1
		case '>':
			positions[i].x += 1
		}
		visited[positions[i]] = true
		i = (i + 1) % actors
	}

	return len(visited)
}

//go:embed input.txt
var puzleInput string

func main() {
	fmt.Printf("Part 1: %d\n", Deliveries(puzleInput, 1))
}
