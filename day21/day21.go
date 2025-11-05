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
