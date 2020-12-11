package main

import (
	"advent"
	"fmt"
	"sort"
)

func main() {
	solution, err := solve1(`input.txt`)
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
