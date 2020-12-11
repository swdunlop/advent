package main

import (
	"advent/test"
	"testing"
)

func TestSolve1(t *testing.T) {
	test.Solve(t, func(path string) (int, error) {
		return solve1(path, 5)
	}, `test.txt`, 127)
}

func TestSolve2(t *testing.T) {
	test.Solve(t, func(path string) (int, error) {
		return solve2(path, 127)
	}, `test.txt`, 62)
}
