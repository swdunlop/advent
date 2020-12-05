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

func solve1(file string) (int, error) {
	at, err := readAtlas(file)
	if err != nil {
		return 0, err
	}
	return at.countTrees(3, 1), nil
}

func solve2(file string) (int, error) {
	at, err := readAtlas(file)
	if err != nil {
		return 0, err
	}
	return at.countTrees(1, 1) * // Right 1, down 1.
			at.countTrees(3, 1) * // Right 3, down 1. (This is the slope you already checked.)
			at.countTrees(5, 1) * // Right 5, down 1.
			at.countTrees(7, 1) * // Right 7, down 1.
			at.countTrees(1, 2), // Right 1, down 2. (I guess it teleports, since this is not an integral)
		nil
}

func readAtlas(file string) (atlas, error) {
	lines, err := advent.ReadLines(file)
	if err != nil {
		return nil, err
	}
	return atlas(lines), nil
}

type atlas []string

func (at atlas) countTrees(right, down int) int {
	row, col, n := 0, 0, 0
	for row < len(at) {
		if at.at(row, col) == tree {
			n++
		}
		row += down
		col += right
	}
	return n
}
func (at atlas) at(row, col int) byte {
	line := at[row]
	return line[col%len(line)]
}

const (
	tree = byte('#')
	open = byte('.')
)
