package main

import (
	"math"
	"regexp"

	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

type Character struct{ hp, damage, armour int }

type Item struct{ cost, damage, armour int }
type Shop map[string]map[string]Item
type ShoppingPlan map[string]struct{ min, max int }

type ShoppingOptions []struct {
	spent        int
	equippedChar Character
}

var charRe = regexp.MustCompile(`^Hit Points: (\d+)\nDamage: (\d+)\nArmor: (\d+)\n$`)

func ParseCharacter(s string) Character {
	c := Character{}
	input.Parse(charRe, s, &c.hp, &c.damage, &c.armour)
	return c
}

func PlayerWins(player, boss Character) bool {
	playerDPR := max(player.damage-boss.armour, 1)
	roundBossDies := math.Ceil(float64(boss.hp) / float64(playerDPR))

	bossDPR := max(boss.damage-player.armour, 1)
	roundPlayerDies := math.Ceil(float64(player.hp) / float64(bossDPR))

	return roundPlayerDies >= roundBossDies
}

func LetsGoShopping(char Character, shop Shop, plan ShoppingPlan) ShoppingOptions {
	options := ShoppingOptions{}

	// for k, v := range plan {

	// }
	return options
}

func Combinations[T any](pool []T, min, max int) [][]T {

	combos := [][]T{}
	for r := min; r <= max; r++ {
		combos = append(combos, combinationsOfLengthR(pool, r)...)
	}
	return combos
}

// I'm missing Python's itertools.combinations().
// Made this function by following its pseudocde:
// https://docs.python.org/3/library/itertools.html#itertools.combinations
// The original relies on some Python feautres, making this version harder to follow.
func combinationsOfLengthR[T any](pool []T, r int) [][]T {
	combos := [][]T{}
	n := len(pool)
	if r > n {
		return combos
	}

	indices := make([]int, r)
	for i := range r {
		indices[i] = i
	}

	getCombo := func() []T {
		c := make([]T, r)
		for i, poolIdx := range indices {
			c[i] = pool[poolIdx]
		}
		return c
	}

	combos = append(combos, getCombo())
	for {
		i := r - 1
		didBreak := false
		for ; i >= 0; i-- {
			if indices[i] != i+n-r {
				didBreak = true
				break
			}
		}
		if !didBreak {
			return combos
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		combos = append(combos, getCombo())
	}
}
