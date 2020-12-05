package main

import (
	"advent"
	"fmt"
)

func main() {
	solution, err := solve2(`input.txt`)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve2(file string) (int, error) {
	input, table, err := load(file)
	if err != nil {
		return 0, err
	}
	for i, n := range input {
		for _, m := range input[i:] {
			d := 2020 - n - m
			if _, ok := table[d]; ok {
				return n * m * d, nil
			}
		}
	}
	return 0, fmt.Errorf(`no solution found in %q`, file)
}

func solve1(file string) (int, error) {
	input, table, err := load(file)
	if err != nil {
		return 0, err
	}
	for _, n := range input {
		d := 2020 - n
		if _, ok := table[d]; ok {
			return n * d, nil
		}
	}
	return 0, fmt.Errorf(`no solution found in %q`, file)
}

func load(file string) (input []int, table map[int]struct{}, err error) {
	input, err = advent.ReadInts(file)
	if err != nil {
		return
	}
	table = make(map[int]struct{}, len(input))
	for _, n := range input {
		table[n] = struct{}{}
	}
	return
}
