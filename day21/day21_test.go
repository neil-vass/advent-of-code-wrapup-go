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
