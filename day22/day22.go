package main

import (
	_ "embed"
	"fmt"
	"maps"
	"regexp"

	"github.com/neil-vass/advent-of-code-2015-go/shared/graph"
	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

//go:embed input.txt
var puzzleData string

func main() {
	boss := Boss{}
	bossRe := regexp.MustCompile(`^Hit Points: (\d+)\nDamage: (\d+)\n$`)
	input.Parse(bossRe, puzzleData, &boss.HP, &boss.Damage)

	game := Game{
		Spellbook: Spellbook{
			"Magic Missile": {Cost: 53, Effect: MagicMissile},
			"Drain":         {Cost: 73, Effect: Drain},
			"Shield":        {Cost: 113, Effect: Shield, Duration: 6},
			"Poison":        {Cost: 173, Effect: Poison, Duration: 6},
			"Recharge":      {Cost: 229, Effect: Recharge, Duration: 5},
		},
		InitialState: GameState{
			Player:       Player{HP: 50, Armour: 0, Mana: 500},
			Boss:         boss,
			ActiveSpells: ActiveSpells{},
		},
		CheapestDamage: 173 / 18.0, // Poison's mana cost per HP
		HardMode:       false,
	}

	fmt.Printf("Part 1: %v\n", Solve(game))
	game.HardMode = true
	fmt.Printf("Part 2: %v\n", Solve(game))
}

func Solve(game Game) int {
	goalFound, cost := graph.A_StarSearch(game, game.InitialState)
	if !goalFound {
		panic("No route to goal!")
	}
	return cost
}

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
		state.ActiveSpells[spellName] = spell.Duration
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

func Drain(state GameState) GameState {
	state.Boss.HP -= 2
	state.Player.HP += 2
	return state
}

func Shield(state GameState) GameState {
	state.Player.Armour = 7
	return state
}

func Poison(state GameState) GameState {
	state.Boss.HP -= 3
	return state
}

func Recharge(state GameState) GameState {
	state.Player.Mana += 101
	return state
}

type Game struct {
	Spellbook      Spellbook
	InitialState   GameState
	CheapestDamage float64
	HardMode       bool
}

func (g Game) Neighbours(node GameState) []graph.NodeCost[GameState] {
	neighbours := []graph.NodeCost[GameState]{}
	for spellName, spell := range g.Spellbook {
		state := GameState{
			Player:       node.Player,
			Boss:         node.Boss,
			ActiveSpells: maps.Clone(node.ActiveSpells),
		}

		valid, newState := g.PlayRound(spellName, state)
		if valid {
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

// Returns (bool valid, newState GameState)
// valid: False if named spell couldn't be cast, or if player dies this round.
// newState: a copy of the given state, updated after the player and boss actions.
func (g Game) PlayRound(spellName string, state GameState) (bool, GameState) {

	// Rule from part 2
	if g.HardMode {
		state.Player.HP -= 1
		if state.Player.HP <= 0 {
			return false, state
		}
	}

	state = g.ApplyEffects(state)

	spellTooExpensive := g.Spellbook[spellName].Cost > state.Player.Mana
	_, SpellAlreadyActive := state.ActiveSpells[spellName]
	if spellTooExpensive || SpellAlreadyActive {
		return false, state
	}

	// Player turn.
	state = g.Spellbook.Cast(spellName, state)

	// Boss turn, if he's alive.
	state = g.ApplyEffects(state)
	if state.Boss.HP > 0 {
		state.Player.HP -= max(state.Boss.Damage-state.Player.Armour, 1)
	}

	// If player died we don't consider this a valid move, like the king in chess.
	return state.Player.HP > 0, state
}

// Apply effects of any long-running spells, decrementing their timers.
func (g Game) ApplyEffects(state GameState) GameState {
	// Wear off any ongoing buffs - they'll reapply if active spells are still up.
	state.Player.Armour = 0

	for spellName, timer := range state.ActiveSpells {
		state = g.Spellbook[spellName].Effect(state)
		if timer > 1 {
			state.ActiveSpells[spellName] -= 1
		} else {
			delete(state.ActiveSpells, spellName)
		}
	}
	return state
}
