package main

import (
	"math"
	"regexp"

	"github.com/neil-vass/advent-of-code-2015-go/shared/input"
)

type Character struct{ hp, damage, armour int }

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
