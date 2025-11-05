package main

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeliveries(t *testing.T) {
	tests := []struct {
		house int
		want  int
	}{
		{1, 10},
		{2, 30},
		{3, 40},
		{4, 70},
		{8, 150},
		{9, 130},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.house), func(t *testing.T) {
			got := Deliveries(tt.house)
			if got != tt.want {
				t.Errorf("Deliveries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorsOf(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{4, []int{1, 2, 4}},
		{8, []int{1, 2, 4, 8}},
		{100, []int{1, 2, 50, 4, 25, 5, 20, 10, 100}},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := FactorsOf(tt.n)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("Contents mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// In part 2, elves deliver (11 * their number) presents to a max of 50 houses.
func TestDeliveriesPart2(t *testing.T) {
	tests := []struct {
		house int
		want  int
	}{
		{1, 11},
		{2, 33},
		{100, (2*11 + 50*11 + 4*11 + 25*11 + 5*11 + 20*11 + 10*11 + 100*11)}, // Elf 1 has been to 50 houses already
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.house), func(t *testing.T) {
			got := DeliveriesPart2(tt.house)
			if got != tt.want {
				t.Errorf("Deliveries() = %v, want %v", got, tt.want)
			}
		})
	}
}
