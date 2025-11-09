package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// I think the longest we can survive is 50 rounds:
// Boss always does at least 1 damage, we have 50 HP.
// If we have both healing (2) and armour (7) on every round,
// we'll be in that best case scenario; otherwie, less than 50 rounds.
// Generating all the "what might I cast in each of 50 rounds" posiblities
// and inspecting them is a huge task: 5^50 is a 35-digit numnber.

// We're looking for the cheapest (least mana) way to end the fight.
// So the ideal case would be: find the spell that has the lowest cost
// per damage point, and cast that every round, while tanking damage ourself
// and never worrying about running out of mana.

// By inspection:
// Magic Missile does 4 damage, costing 13.25 mana per damage point.
// Drain does 2 damage, costing 36.5 mana per point.
// Poison does 18 damage (eventually), costing 9.61 mana per point.
// In an ideal world I'd just use Poison till the boss dies.

// This is sounding like A*:
// Our goal is boss at zero HP
// Our cost to minimize is mana points
// Our heuristic (best case from here to there) is boss's remaining HP * 9.61.
// Our "neighbours" (possible next steps) are the spells we have mana for,
// as long as we have HP. If we die or can't afford spells this route is a dead end.

func TestSpellcasting(t *testing.T) {
	spellbook := Spellbook{
		"Magic Missile": {Cost: 53, Cast: MagicMissile},
	}

	state := GameState{
		Player: Player{HP: 50, Armour: 0, Mana: 500},
		Boss:   Boss{HP: 16, Damage: 10},
	}

	got := spellbook["Magic Missile"].Cast(state)

	want := GameState{
		Player: Player{HP: 50, Armour: 0, Mana: 447},
		Boss:   Boss{HP: 12, Damage: 10},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart1(t *testing.T) {
	game := Game{
		Spellbook: Spellbook{
			"Magic Missile": {Cost: 53, Cast: MagicMissile},
		},
		InitialState: GameState{
			Player: Player{HP: 50, Armour: 0, Mana: 500},
			Boss:   Boss{HP: 16, Damage: 10},
		},
		CheapestDamage: 13.25,
	}

	got := SolvePart1(game)
	want := 53 * 4 // Cast Magic Missile 4 times and win!
	if got != want {
		t.Errorf("SolvePart1()=%v, want %v", got, want)
	}
}
