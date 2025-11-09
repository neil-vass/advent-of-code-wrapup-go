package main

import "github.com/neil-vass/advent-of-code-2015-go/shared/graph"

type Spell struct {
	Cost   int
	Cast   SpellcastFn
	Effect SpellEffectFn
}
type Spellbook map[string]Spell

type Player struct{ HP, Armour, Mana int }
type Boss struct{ HP, Damage int }
type ActiveSpells map[string]int

type GameState struct {
	Player Player
	Boss   Boss
}

type SpellcastFn func(state GameState) GameState
type SpellEffectFn func(state GameState) GameState

func MagicMissile(state GameState) GameState {
	state.Player.Mana -= 53
	state.Boss.HP -= 4
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
			newState := g.Spellbook[name].Cast(node)
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
