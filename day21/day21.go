package main

import (
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
