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
	fmt.Printf("Part 1: %d\n", Find(target))
}

func Find(target int) int {
	i := 100_000
	for {
		if i%100_000 == 0 {
			fmt.Println("Testing", i)
		}

		if Deliveries(i) >= target {
			return i
		}

		i++
	}
}
func Deliveries(house int) int {
	var count int
	for _, factor := range FactorsOf(house) {
		count += 10 * factor
	}
	return count
}

func FactorsOf(house int) []int {
	factors := []int{1}
	for n := 2; n*n <= house; n++ {
		if house%n == 0 {
			factors = append(factors, n)
			if n != house/n {
				factors = append(factors, house/n)
			}
		}
	}
	if house != 1 {
		factors = append(factors, house)
	}
	return factors
}
