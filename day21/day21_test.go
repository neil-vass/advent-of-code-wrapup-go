package main

import (
	"testing"
)

func TestParseCharacter(t *testing.T) {
	input := "Hit Points: 10\nDamage: 2\nArmor: 1\n"
	got := ParseCharacter(input)
	want := Character{hp: 10, damage: 2, armour: 1}
	if got != want {
		t.Errorf("ParseCharacter()=%v, want %v", got, want)
	}
}

func TestBossBattle(t *testing.T) {
	player := Character{hp: 8, damage: 5, armour: 5}
	boss := Character{hp: 12, damage: 7, armour: 2}

	if !PlayerWins(player, boss) {
		t.Error("Expected player to win, but they lost!")
	}
}
