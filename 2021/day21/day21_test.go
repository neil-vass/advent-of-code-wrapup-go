package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2015-go/shared/assert"
)

func TestGameOutcomes(t *testing.T) {
	t.Run("Just one roll", func(t *testing.T) {
		game := GameState{
			P1: Player{Pos: 4, Needs: 1},
			P2: Player{Pos: 8, Needs: 1},
		}

		p1Wins, p2Wins := GameOutcomes(game)
		assert.Equal(t, p1Wins, 27)
		assert.Equal(t, p2Wins, 0)
	})

	t.Run("Solves example", func(t *testing.T) {
		game := GameState{
			P1: Player{Pos: 4, Needs: 21},
			P2: Player{Pos: 8, Needs: 21},
		}

		p1Wins, p2Wins := GameOutcomes(game)
		assert.Equal(t, p1Wins, 444356092776315)
		assert.Equal(t, p2Wins, 341960390180808)
	})
}
