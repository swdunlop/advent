package main

import (
	"advent/test"
	"testing"
)

func TestSolve1(t *testing.T) { test.Solve(t, solve1, `test1.txt`, 2) }

func TestSolve2(t *testing.T) { test.Solve(t, solve2, `test2.txt`, 4) }
