package main

import (
	"advent/test"
	"testing"
)

func TestSolve1(t *testing.T) {
	test.Solve(t, solve1, `test1.txt`, 35)
	test.Solve(t, solve1, `test2.txt`, 220)
}

// func TestSolve2(t *testing.T) { test.Solve(t, solve2, `test2.txt`, 126) }
