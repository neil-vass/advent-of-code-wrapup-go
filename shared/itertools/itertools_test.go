package itertools

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		name   string
		pool   []int
		length int
		want   [][]int
	}{
		{name: "2 out of 3", pool: []int{1, 2, 3}, length: 2, want: [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{name: "1 out of 3", pool: []int{1, 2, 3}, length: 1, want: [][]int{{1}, {2}, {3}}},
		{name: "Zero length combos", pool: []int{1, 2, 3}, length: 0, want: [][]int{{}}}, // A single zero-size result.
		{name: "Oversize combos", pool: []int{1, 2, 3}, length: 5, want: [][]int{}},      // No result, can't make any.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combinations(tt.pool, tt.length)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("Contents mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	weapons := [][]string{{"Dagger"}, {"Shortsword"}}
	armour := [][]string{{}, {"Leather"}}

	got := Product(weapons, armour)
	want := [][][]string{
		{{"Dagger"}, {}},
		{{"Dagger"}, {"Leather"}},
		{{"Shortsword"}, {}},
		{{"Shortsword"}, {"Leather"}},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}
