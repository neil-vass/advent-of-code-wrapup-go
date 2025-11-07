package main

import (
	_ "embed"
	"fmt"
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

//go:embed input.txt
var puzzleData string

func main() {
	player := Character{HP: 100, Damage: 0, Armour: 0}
	boss := ParseCharacter(puzzleData)

	shop := Shop{
		"Weapons": {
			"Dagger":     {Cost: 8, Damage: 4, Armour: 0},
			"Shortsword": {10, 5, 0},
			"Warhammer":  {25, 6, 0},
			"Longsword":  {40, 7, 0},
			"Greataxe":   {74, 8, 0},
		},
		"Armour": {
			"Leather":    {13, 0, 1},
			"Chainmail":  {31, 0, 2},
			"Splintmail": {53, 0, 3},
			"Bandedmail": {75, 0, 4},
			"Platemail":  {102, 0, 5},
		},
		"Rings": {
			"Damage +1":  {25, 1, 0},
			"Damage +2":  {50, 2, 0},
			"Damage +3":  {100, 3, 0},
			"Defense +1": {20, 0, 1},
			"Defense +2": {40, 0, 2},
			"Defense +3": {80, 0, 3},
		},
	}

	plan := ShoppingPlan{
		"Weapons": {Min: 1, Max: 1},
		"Armour":  {Min: 0, Max: 1},
		"Rings":   {Min: 0, Max: 2},
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(player, boss, shop, plan))
}

func SolvePart1(player, boss Character, shop Shop, plan ShoppingPlan) int {
	options := LetsGoShopping(player, shop, plan)
	return CheapestWayToWin(options, boss)
}

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

func CheapestWayToWin(options ShoppingOptions, boss Character) int {
	cheapestSoFar := math.MaxInt
	for _, opt := range options {
		if PlayerWins(opt.EquippedChar, boss) && opt.Spent < cheapestSoFar {
			cheapestSoFar = opt.Spent
		}
	}
	return cheapestSoFar
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
