package main

import (
	"testing"
)

type tableTest[V comparable] struct {
	name string // description of this test case
	got  V
	want V
}

func assertTableTests[V comparable](t *testing.T, tests []tableTest[V]) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			if tt.got != tt.want {
				t.Errorf("%v: got %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestDeliveries(t *testing.T) {
	tests := []tableTest[int]{
		{name: "One step", got: Deliveries(">", 1), want: 2},
		{name: "Square", got: Deliveries(">v<", 1), want: 4},
		{name: "Robo-Santa", got: Deliveries("^v^v^v^v^v", 2), want: 11},
	}

	assertTableTests(t, tests)
}
