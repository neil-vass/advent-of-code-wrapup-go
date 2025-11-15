package main

import (
	"slices"
	"testing"

	"github.com/neil-vass/advent-of-code-2015-go/shared/assert"
)

func TestParsePresent(t *testing.T) {
	assert.Equal(t, ParsePresent("2x3x4"), Present{2, 3, 4})
}

func TestPresent_PaperNeeded(t *testing.T) {
	assert.Equal(t, PaperNeeded(Present{2, 3, 4}), 58)
	assert.Equal(t, PaperNeeded(Present{1, 1, 10}), 43)
}

func TestSolver(t *testing.T) {
	lines := slices.Values([]string{"2x3x4", "1x1x10"})
	got := Solve(lines, PaperNeeded)
	assert.Equal(t, got, 58+43)
}

func TestPresent_RibbonNeeded(t *testing.T) {
	assert.Equal(t, RibbonNeeded(Present{2, 3, 4}), 34)
	assert.Equal(t, RibbonNeeded(Present{1, 1, 10}), 14)
}
