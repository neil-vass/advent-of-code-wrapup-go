package main

import (
	"maps"
	"math"
	"regexp"
	"slices"

	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

type Character struct{ HP, Damage, Armour int }

type Item struct{ Cost, Damage, Armour int }
type Shop map[string]map[string]Item
type ShoppingPlan map[string]struct{ Min, Max int }

type Foo struct {
	Spent        int
	EquippedChar Character
}
type ShoppingOptions []Foo

var charRe = regexp.MustCompile(`^Hit Points: (\d+)\nDamage: (\d+)\nArmor: (\d+)\n$`)

func ParseCharacter(s string) Character {
	c := Character{}
	input.Parse(charRe, s, &c.HP, &c.Damage, &c.Armour)
	return c
}

func PlayerWins(player, boss Character) bool {
	playerDPR := max(player.Damage-boss.Armour, 1)
	roundBossDies := math.Ceil(float64(boss.HP) / float64(playerDPR))

	bossDPR := max(boss.Damage-player.Armour, 1)
	roundPlayerDies := math.Ceil(float64(player.HP) / float64(bossDPR))

	return roundPlayerDies >= roundBossDies
}

func LetsGoShopping(char Character, shop Shop, plan ShoppingPlan) ShoppingOptions {
	options := ShoppingOptions{}

	collector := [][][]Item{}
	for k, v := range shop {
		values := slices.Collect(maps.Values(v))
		optionsForThisKey := Combinations(values, plan[k].Min, plan[k].Max)
		collector = append(collector, optionsForThisKey)
	}

	allPurchasePlans := Product(collector...)

	for _, p := range allPurchasePlans {
		equippedChar := char
		spent := 0
		for _, itemsOfOneType := range p {
			for _, item := range itemsOfOneType {
				spent += item.Cost
				equippedChar.Damage += item.Damage
				equippedChar.Armour += item.Armour
			}
		}
		options = append(options, Foo{Spent: spent, EquippedChar: equippedChar})
	}

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

// I'm missing Python's itertools.combinations().
// Made this function by following its pseudocde:
// https://docs.python.org/3.7/library/itertools.html#itertools.product
func Product[T any](pools ...[]T) [][]T {
	result := [][]T{{}}
	for _, pool := range pools {
		updatedResult := [][]T{}
		for _, x := range result {
			for _, y := range pool {
				updatedResult = append(updatedResult, append(x, y))
			}
		}
		result = updatedResult
	}
	return result
}
