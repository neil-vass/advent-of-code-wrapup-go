package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2015-go/shared/assert"
)

func TestPositionInSequence(t *testing.T) {
	assert.Equal(t, PositionInSequence(1, 1), 1)
	assert.Equal(t, PositionInSequence(2, 1), 2)
	assert.Equal(t, PositionInSequence(3, 1), 4)
	assert.Equal(t, PositionInSequence(4, 1), 7)
	assert.Equal(t, PositionInSequence(5, 1), 11)
	assert.Equal(t, PositionInSequence(6, 1), 16)

	assert.Equal(t, PositionInSequence(1, 2), 3)
	assert.Equal(t, PositionInSequence(2, 2), 5)

	assert.Equal(t, PositionInSequence(4, 2), 12)
	assert.Equal(t, PositionInSequence(1, 5), 15)
	assert.Equal(t, PositionInSequence(4, 3), 18)
}

func TestCodeFor(t *testing.T) {
	assert.Equal(t, CodeFor(1), 20151125)
	assert.Equal(t, CodeFor(2), 31916031)
	assert.Equal(t, CodeFor(10), 30943339)
}
