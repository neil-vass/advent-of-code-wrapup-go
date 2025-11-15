package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2015-go/shared/itertools"
)

//go:embed input.txt
var puzzleData string

func main() {
	lines := strings.Split(strings.TrimSpace(puzzleData), "\n")
	packages := make([]int, len(lines))
	for i, ln := range lines {
		p, _ := strconv.Atoi(ln)
		packages[i] = p
	}

	if ok, result := MinQuantumScore(packages, 3); ok {
		fmt.Printf("Part 1: %v\n", result)
	} else {
		fmt.Printf("Part 1: failed to find solution\n")
	}

	if ok, result := MinQuantumScore(packages, 4); ok {
		fmt.Printf("Part 2: %v\n", result)
	} else {
		fmt.Printf("Part 2: failed to find solution\n")
	}

}

func TargetWeight(packages []int, numGroups int) int {
	total := 0
	for _, w := range packages {
		total += w
	}
	return total / numGroups
}

func MinQuantumScore(packages []int, numGroups int) (bool, int) {
	targetWeight := TargetWeight(packages, numGroups)

	for groupSize := 1; groupSize < len(packages)-2; groupSize++ {
		quantumScores := []int{}
		for _, group := range itertools.Combinations(packages, groupSize) {

			if meetsTarget(group, targetWeight) && canBalance(packages, group, targetWeight, numGroups-1) {
				quantumScores = append(quantumScores, score(group))
			}

		}
		if len(quantumScores) > 0 {
			return true, slices.Min(quantumScores)
		}
	}
	return false, -1
}

func meetsTarget(group []int, targetWeight int) bool {
	weight := 0
	for _, w := range group {
		weight += w
	}
	return weight == targetWeight
}

func canBalance(packages []int, firstGroup []int, targetWeight int, numGroupsRemaining int) bool {
	usedSet := map[int]bool{}
	for _, p := range firstGroup {
		usedSet[p] = true
	}

	remainder := []int{}
	for _, p := range packages {
		if _, alreadyUsed := usedSet[p]; !alreadyUsed {
			remainder = append(remainder, p)
		}
	}

	for groupSize := 1; groupSize < len(remainder)-1; groupSize++ {
		for _, group := range itertools.Combinations(packages, groupSize) {
			if meetsTarget(group, targetWeight) {
				if numGroupsRemaining <= 2 {
					return true
				}
				return canBalance(remainder, group, targetWeight, numGroupsRemaining-1)
			}
		}
	}
	return false
}

func score(group []int) int {
	quantumScore := 1
	for _, w := range group {
		quantumScore *= w
	}
	return quantumScore
}
