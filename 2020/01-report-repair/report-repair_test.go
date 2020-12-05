package main

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	test(t, solve1, "test.txt", 514579)
}

func TestSolve2(t *testing.T) {
	test(t, solve2, "test.txt", 241861950)
}

func test(t *testing.T, try func(string) (int, error), file string, x int) {
	s, err := try(file)
	if err != nil {
		t.Fatal(err)
	}
	if s != x {
		t.Fatalf(`expected %v, got %v`, x, s)
	}
}
