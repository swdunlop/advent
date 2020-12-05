package main

import (
	"advent/test"
	"testing"
)

func TestSolve1(t *testing.T) { test.Solve(t, solve1, `test.txt`, 7) }

func TestSolve2(t *testing.T) { test.Solve(t, solve2, `test.txt`, 336) }
