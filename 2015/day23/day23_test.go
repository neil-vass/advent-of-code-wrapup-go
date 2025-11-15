package main

import (
	"testing"
)

func TestComputer_Run(t *testing.T) {

	c := NewComputer(
		"inc a",
		"jio a, +2",
		"tpl a",
		"inc a")
	c.Run()

	got := c.registers["a"]
	want := 2
	if got != want {
		t.Errorf("Register a = %v, want %v", got, want)
	}
}
