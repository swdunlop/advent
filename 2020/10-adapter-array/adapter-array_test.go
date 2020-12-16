package main

import (
	"advent/test"
	"testing"
)

func TestSolve1(t *testing.T) {
	test.Solve(t, solve1, `test1.txt`, 35)
	test.Solve(t, solve1, `test2.txt`, 220)
}

func TestSolve2(t *testing.T) {
	test.Solve(t, func(path string) (int, error) {
		return solve2(path, 22)
	}, `test1.txt`, 8)

	test.Solve(t, func(path string) (int, error) {
		return solve2(path, 52)
	}, `test2.txt`, 19208)
}
