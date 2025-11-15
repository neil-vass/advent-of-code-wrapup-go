package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParseCharacter(t *testing.T) {
	input := "Hit Points: 10\nDamage: 2\nArmor: 1\n"
	got := ParseCharacter(input)
	want := Character{HP: 10, Damage: 2, Armour: 1}
	if got != want {
		t.Errorf("ParseCharacter()=%v, want %v", got, want)
	}
}

func TestPlayerWins(t *testing.T) {
	player := Character{HP: 8, Damage: 5, Armour: 5}
	boss := Character{HP: 12, Damage: 7, Armour: 2}

	if !PlayerWins(player, boss) {
		t.Error("Expected player to win, but they lost!")
	}
}

func fakeShop() Shop {
	return Shop{
		"Weapons": {
			"Dagger":     {Cost: 8, Damage: 4, Armour: 0},
			"Shortsword": {Cost: 10, Damage: 5, Armour: 0},
		},
		"Armour": {
			"Leather": {Cost: 13, Damage: 0, Armour: 1},
		},
	}
}

func fakeShoppingPlan() ShoppingPlan {
	return ShoppingPlan{
		"Weapons": {Min: 1, Max: 1},
		"Armour":  {Min: 0, Max: 1},
	}
}

func TestShopping(t *testing.T) {
	player := Character{HP: 1, Damage: 0, Armour: 0}
	got := LetsGoShopping(player, fakeShop(), fakeShoppingPlan())

	want := []ShoppingOption{
		{Spent: 8, EquippedChar: Character{1, 4, 0}},
		{Spent: 21, EquippedChar: Character{1, 4, 1}},
		{Spent: 10, EquippedChar: Character{1, 5, 0}},
		{Spent: 23, EquippedChar: Character{1, 5, 1}},
	}

	// Order got and want for comparison.
	// This works because every entry has a different Spent.
	less := func(a, b ShoppingOption) bool { return a.Spent < b.Spent }

	diff := cmp.Diff(want, got, cmpopts.SortSlices(less))
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart1(t *testing.T) {
	player := Character{HP: 1, Damage: 0, Armour: 0}
	boss := Character{HP: 5, Damage: 100, Armour: 0}

	got := SolvePart1(player, boss, fakeShop(), fakeShoppingPlan())
	want := 10 // Shortsword means you win in one hit, armour not needed.
	if got != want {
		t.Errorf("SolvePart1()=%v, want %v", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	player := Character{HP: 1, Damage: 0, Armour: 0}
	boss := Character{HP: 5, Damage: 100, Armour: 0}

	got := SolvePart2(player, boss, fakeShop(), fakeShoppingPlan())
	want := 21 // Dagger and Leather mean you still die in the first round
	if got != want {
		t.Errorf("SolvePart1()=%v, want %v", got, want)
	}
}

func TestCombinationsOfVaryingLength(t *testing.T) {
	pool := []int{0, 1, 2}
	min := 0
	max := 2

	got := CombinationsOfVaryingLength(pool, min, max)
	want := [][]int{{}, {0}, {1}, {2}, {0, 1}, {0, 2}, {1, 2}}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}
