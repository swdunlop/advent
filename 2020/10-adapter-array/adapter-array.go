package main

import (
	"advent"
	"fmt"
	"sort"
)

func main() {
	solution, err := solve2(`input.txt`, 168+3)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve1(file string) (int, error) {
	adapters, err := advent.ReadInts(file)
	if err != nil {
		return 0, err
	}
	sort.Ints(adapters)

	var ones, threes, prev int
	for _, next := range adapters {
		switch next - prev {
		case 0, 2:
		case 1:
			ones++
		case 3:
			threes++
		default:
			break
		}

		prev = next
	}
	threes++ // include the device

	return ones * threes, nil
}

func solve2(file string, target int) (int, error) {
	adapters, err := advent.ReadInts(file)
	if err != nil {
		return 0, err
	}
	sort.Ints(adapters)

	t := make(table, len(adapters))
	t[0] = entry{}
	for _, adapter := range adapters {
		t[adapter] = entry{}
	}
	return t.walk(0, target), nil
}

type table map[int]entry

func (t table) walk(pos, target int) int {
	if pos == target {
		return 1
	}

	e, ok := t[pos]
	if !ok {
		return 0
	}
	if e.walked {
		return e.paths
	}

	paths := t.walk(pos+1, target) + t.walk(pos+2, target) + t.walk(pos+3, target)
	t[pos] = entry{true, paths}
	return paths
}

type entry struct {
	walked bool
	paths  int
}
