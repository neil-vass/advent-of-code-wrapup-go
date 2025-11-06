package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestShopping(t *testing.T) {
	player := Character{HP: 100, Damage: 0, Armour: 0}
	shop := Shop{
		"Weapons": {
			"Dagger":     {Cost: 8, Damage: 4, Armour: 0},
			"Shortsword": {Cost: 10, Damage: 5, Armour: 0},
		},
		"Armour": {
			"Leather": {Cost: 13, Damage: 0, Armour: 1},
		},
	}
	plan := ShoppingPlan{
		"Weapons": {Min: 1, Max: 1},
		"Armour":  {Min: 0, Max: 1},
	}
	got := LetsGoShopping(player, shop, plan)

	want := ShoppingOptions{
		{Spent: 8, EquippedChar: Character{100, 4, 0}},
		{Spent: 21, EquippedChar: Character{100, 4, 1}},
		{Spent: 10, EquippedChar: Character{100, 5, 0}},
		{Spent: 23, EquippedChar: Character{100, 5, 1}},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestCombinations(t *testing.T) {
	pool := []int{0, 1, 2}
	min := 0
	max := 2

	got := Combinations(pool, min, max)
	want := [][]int{{}, {0}, {1}, {2}, {0, 1}, {0, 2}, {1, 2}}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestProduct(t *testing.T) {
	weapons := [][]string{{"Dagger"}, {"Shortsword"}}
	armour := [][]string{{}, {"Leather"}}

	got := Product(weapons, armour)
	want := [][][]string{
		{{"Dagger"}, {}},
		{{"Dagger"}, {"Leather"}},
		{{"Shortsword"}, {}},
		{{"Shortsword"}, {"Leather"}},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}
