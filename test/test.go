package test

import "testing"

// Solve tries a solution with the provided input string, expecting x as a solution.
func Solve(t *testing.T, try func(string) (int, error), file string, x int) {
	s, err := try(file)
	if err != nil {
		t.Fatal(err)
	}
	if s != x {
		t.Fatalf(`expected %v, got %v`, x, s)
	}
}
