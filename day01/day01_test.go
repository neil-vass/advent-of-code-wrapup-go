package main

import "testing"

func assertEqual[V comparable](t *testing.T, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindFloor(t *testing.T) {
	assertEqual(t, FindFloor("(())"), 0)
	assertEqual(t, FindFloor("(()(()("), 3)
}

func TestPosForBasement(t *testing.T) {
	assertEqual(t, PosForBasement(")"), 1)
	assertEqual(t, PosForBasement("()())"), 5)
}
