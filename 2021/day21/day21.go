package main

import (
	_ "embed"
	"fmt"
	"regexp"

	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

//go:embed input.txt
var puzzleData string

func main() {
	var state GameState
	inputRe := regexp.MustCompile(`^Player 1 starting position: (\d+)\nPlayer 2 starting position: (\d+)\n$`)
	err := input.Parse(inputRe, puzzleData, &state.P1.Pos, &state.P2.Pos)
	if err != nil {
		panic(err)
	}
	state.P1.Needs = 21
	state.P2.Needs = 21

	p1Wins, p2Wins := GameOutcomes(state)
	winner := max(p1Wins, p2Wins)
	fmt.Printf("Part 2: %v\n", winner)
}

var dice = []int{1, 2, 3}
var trackLength = 10

type Player struct{ Pos, Needs int }
type GameState struct {
	P1, P2              Player
	Roll1, Roll2, Roll3 int
	P1TurnDone          bool
}

var cache = map[GameState]struct{ p1Wins, p2Wins uint64 }{}

func GameOutcomes(state GameState) (p1Wins, p2Wins uint64) {
	if res, ok := cache[state]; ok {
		return res.p1Wins, res.p2Wins
	}

	p1Wins, p2Wins = gameOutcomes(state)
	cache[state] = struct {
		p1Wins uint64
		p2Wins uint64
	}{p1Wins, p2Wins}
	return
}

func gameOutcomes(state GameState) (p1Wins, p2Wins uint64) {

	if state.Roll1 == 0 {

		for _, roll := range dice {
			state.Roll1 = roll
			p1WinsOnThisPath, p2WinsOnThisPath := GameOutcomes(state)
			p1Wins += p1WinsOnThisPath
			p2Wins += p2WinsOnThisPath
		}
		return
	}

	if state.Roll2 == 0 {
		for _, roll := range dice {
			state.Roll2 = roll
			p1WinsOnThisPath, p2WinsOnThisPath := GameOutcomes(state)
			p1Wins += p1WinsOnThisPath
			p2Wins += p2WinsOnThisPath
		}
		return
	}

	if state.Roll3 == 0 {
		for _, roll := range dice {
			state.Roll3 = roll
			p1WinsOnThisPath, p2WinsOnThisPath := GameOutcomes(state)
			p1Wins += p1WinsOnThisPath
			p2Wins += p2WinsOnThisPath
			cache[state] = struct {
				p1Wins uint64
				p2Wins uint64
			}{p1WinsOnThisPath, p2WinsOnThisPath}
		}
		return
	}

	totalRoll := state.Roll1 + state.Roll2 + state.Roll3

	if !state.P1TurnDone {
		updatedP1 := moveAndScore(state.P1, totalRoll)

		if updatedP1.Needs <= 0 {
			p1Wins++
			return
		}

		// Let's give P2 a go.
		updatedState := GameState{
			P1:         updatedP1,
			P2:         state.P2,
			P1TurnDone: true,
		}
		p1WinsOnThisPath, p2WinsOnThisPath := GameOutcomes(updatedState)
		p1Wins += p1WinsOnThisPath
		p2Wins += p2WinsOnThisPath
		return
	}

	// If you get here: we have a full set of rolls, and P1's had a turn.
	updatedP2 := moveAndScore(state.P2, totalRoll)

	if updatedP2.Needs <= 0 {
		p2Wins++
		return
	}

	updatedState := GameState{
		P1: state.P1,
		P2: updatedP2,
	}
	p1WinsOnThisPath, p2WinsOnThisPath := GameOutcomes(updatedState)
	p1Wins += p1WinsOnThisPath
	p2Wins += p2WinsOnThisPath
	return
}

func moveAndScore(player Player, roll int) Player {
	var updated Player
	updated.Pos = (player.Pos + roll) % trackLength
	if updated.Pos == 0 {
		updated.Pos = trackLength
	}
	updated.Needs = player.Needs - updated.Pos
	return updated
}
