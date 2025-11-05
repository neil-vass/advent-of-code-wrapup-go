package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCharacter(t *testing.T) {
	input := "Hit Points: 10\nDamage: 2\nArmor: 1\n"
	got := ParseCharacter(input)
	want := Character{hp: 10, damage: 2, armour: 1}
	if got != want {
		t.Errorf("ParseCharacter()=%v, want %v", got, want)
	}
}

func TestPlayerWins(t *testing.T) {
	player := Character{hp: 8, damage: 5, armour: 5}
	boss := Character{hp: 12, damage: 7, armour: 2}

	if !PlayerWins(player, boss) {
		t.Error("Expected player to win, but they lost!")
	}
}

func TestShopping(t *testing.T) {
	player := Character{hp: 100, damage: 0, armour: 0}
	shop := Shop{
		"Weapons": {
			"Dagger":     {cost: 8, damage: 4, armour: 0},
			"Shortsword": {cost: 10, damage: 5, armour: 0},
		},
		"Armour": {
			"Leather": {cost: 13, damage: 0, armour: 1},
		},
	}
	plan := ShoppingPlan{
		"Weapons": {min: 1, max: 1},
		"Armour":  {min: 0, max: 1},
	}

	got := LetsGoShopping(player, shop, plan)
	want := ShoppingOptions{
		{spent: 8, equippedChar: Character{100, 4, 0}},
		{spent: 21, equippedChar: Character{100, 4, 1}},
		{spent: 10, equippedChar: Character{100, 5, 0}},
		{spent: 23, equippedChar: Character{100, 5, 1}},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestCombinations(t *testing.T) {
	li := []int{0, 1, 2}
	min := 0
	max := 2

	got := Combinations(li, min, max)
	want := [][]int{{}, {0}, {1}, {2}, {0, 1}, {0, 2}, {1, 2}}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func Combinations[T any](pool []T, min, max int) [][]T {

	combos := [][]T{}
	for r := min; r <= max; r++ {
		combos = append(combos, combinationsOfLengthR(pool, r)...)
	}
	return combos
}

// I'm missing Python's itertools.combinations().
// Made this function by following its pseudocde:
// https://docs.python.org/3/library/itertools.html#itertools.combinations
// The original relies on some Python feautres, making this version harder to follow.
func combinationsOfLengthR[T any](pool []T, r int) [][]T {
	combos := [][]T{}
	n := len(pool)
	if r > n {
		return combos
	}

	indices := make([]int, r)
	for i := range r {
		indices[i] = i
	}

	getCombo := func() []T {
		c := make([]T, r)
		for i, poolIdx := range indices {
			c[i] = pool[poolIdx]
		}
		return c
	}

	combos = append(combos, getCombo())
	for {
		i := r - 1
		didBreak := false
		for ; i >= 0; i-- {
			if indices[i] != i+n-r {
				didBreak = true
				break
			}
		}
		if !didBreak {
			return combos
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		combos = append(combos, getCombo())
	}
}
