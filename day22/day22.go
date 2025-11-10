package main

import (
	"github.com/neil-vass/advent-of-code-2015-go/shared/graph"
)

type Spell struct {
	Cost     int
	Effect   SpellFn
	Duration int
}
type Spellbook map[string]Spell

func (book Spellbook) Cast(spellName string, state GameState) GameState {
	spell := book[spellName]
	state.Player.Mana -= spell.Cost
	if spell.Duration == 0 {
		state = spell.Effect(state)
	} else {
		// add spell name and duration to active spells
	}
	return state
}

type Player struct{ HP, Armour, Mana int }
type Boss struct{ HP, Damage int }
type ActiveSpells map[string]int

type GameState struct {
	Player       Player
	Boss         Boss
	ActiveSpells ActiveSpells
}

type SpellFn func(state GameState) GameState

func MagicMissile(state GameState) GameState {
	state.Boss.HP -= 4
	return state
}

func Shield(state GameState) GameState {
	state.Player.Armour = 7
	return state
}

type Game struct {
	Spellbook      Spellbook
	InitialState   GameState
	CheapestDamage float64
}

// TODO: Currently player just gets to keep firing spells, boss never gets a turn.
func (g Game) Neighbours(node GameState) []graph.NodeCost[GameState] {
	neighbours := []graph.NodeCost[GameState]{}
	for name, spell := range g.Spellbook {
		if spell.Cost <= node.Player.Mana {
			newState := g.Spellbook[name].Effect(node)
			n := graph.NodeCost[GameState]{Node: newState, Cost: spell.Cost}
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func (g Game) Heuristic(from GameState) float64 {
	return float64(from.Boss.HP) * g.CheapestDamage
}

func (g Game) GoalReached(candidate GameState) bool {
	return candidate.Boss.HP <= 0
}

func SolvePart1(game Game) int {
	goalFound, cost := graph.A_StarSearch(game, game.InitialState)
	if !goalFound {
		panic("No route to goal!")
	}
	return cost
}
