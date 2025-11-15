package main

import (
	"testing"
)

func TestTargetWeight(t *testing.T) {
	packages := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	got := TargetWeight(packages, 3)
	want := 20
	if got != want {
		t.Errorf("TargetWeight()=%v, want %v", got, want)
	}
}

func TestMinQuantumScore(t *testing.T) {
	packages := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	found, got := MinQuantumScore(packages, 3)
	want := 99
	if !found {
		t.Errorf("Failed to find a solution")
	}
	if got != want {
		t.Errorf("MinQuantumScore()=%v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	packages := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	found, got := MinQuantumScore(packages, 4)
	want := 44
	if !found {
		t.Errorf("Failed to find a solution")
	}
	if got != want {
		t.Errorf("MinQuantumScore()=%v, want %v", got, want)
	}
}
