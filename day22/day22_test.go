package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/neil-vass/advent-of-code-2015-go/shared/assert"
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

func SimpleGameSetup() Game {
	return Game{
		Spellbook: Spellbook{
			"Magic Missile": {Cost: 53, Effect: MagicMissile},
		},
		InitialState: GameState{
			Player:       Player{HP: 14, Armour: 0, Mana: 500},
			Boss:         Boss{HP: 12, Damage: 1},
			ActiveSpells: ActiveSpells{},
		},
		CheapestDamage: 53 / 4.0, // Magic Missile
	}
}

func TestSpellcasting(t *testing.T) {
	game := SimpleGameSetup()

	got := game.Spellbook.Cast("Magic Missile", game.InitialState)

	want := GameState{
		Player:       Player{HP: 14, Armour: 0, Mana: 447},
		Boss:         Boss{HP: 8, Damage: 1},
		ActiveSpells: ActiveSpells{},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestCanWinManually(t *testing.T) {
	game := SimpleGameSetup()
	game.InitialState.Boss.Damage = 7 // Strong boss, but a shield will save you
	game.Spellbook["Shield"] = Spell{Cost: 113, Effect: Shield, Duration: 6}

	valid, state := game.PlayRound("Shield", game.InitialState)
	assert.Equal(t, valid, true)
	assert.Equal(t, state.Player.HP, 13) // Just 1HP damage for 3 rounds (6 turns)

	valid, state = game.PlayRound("Magic Missile", state)
	assert.Equal(t, valid, true)
	assert.Equal(t, state.Player.HP, 12)
	assert.Equal(t, state.Boss.HP, 8)

	valid, state = game.PlayRound("Magic Missile", state)
	assert.Equal(t, valid, true)
	assert.Equal(t, state.Player.HP, 11)
	assert.Equal(t, state.Boss.HP, 4)

	valid, state = game.PlayRound("Magic Missile", state)
	assert.Equal(t, valid, true)
	assert.Equal(t, state.Player.HP, 11) // Boss dies on player's turn, no damage done
	assert.Equal(t, state.Boss.HP, 0)
}

func TestSolvePart1(t *testing.T) {
	t.Run("Solve simple game", func(t *testing.T) {
		game := SimpleGameSetup()

		got := SolvePart1(game)
		want := 53 * 3 // Cast Magic Missile 3 times and win!
		if got != want {
			t.Errorf("SolvePart1()=%v, want %v", got, want)
		}
	})

	t.Run("Recognize unwinnable game", func(t *testing.T) {
		defer func() {
			got := fmt.Sprint(recover())
			want := "No route to goal!"
			if got != want {
				t.Errorf("Wrong panic reason: got %#v, want %#v", got, want)
			}
		}()

		game := SimpleGameSetup()
		game.InitialState.Boss.Damage = 7 // Strong boss will win
		SolvePart1(game)
	})

	t.Run("Choose second spell when needed", func(t *testing.T) {
		game := SimpleGameSetup()
		game.InitialState.Boss.Damage = 7 // Strong boss, but a shield will save you
		game.Spellbook["Shield"] = Spell{Cost: 113, Effect: Shield, Duration: 6}

		got := SolvePart1(game)
		want := 113 + 53*3 // Shields up, cast Magic Missile 3 times and win!
		if got != want {
			t.Errorf("SolvePart1()=%v, want %v", got, want)
		}
	})
}
