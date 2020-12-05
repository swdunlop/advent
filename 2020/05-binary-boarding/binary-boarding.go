package main

import (
	"advent"
	"fmt"
	"sort"
)

func main() {
	solution, err := solve2(`input.txt`)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve2(input string) (int, error) {
	paths, err := advent.ReadLines(input)
	if err != nil {
		return 0, err
	}

	seats := make([]int, len(paths))
	for i, path := range paths {
		seat, err := locateSeat(path, 128, 8)
		if err != nil {
			return 0, err
		}
		seats[i] = seat
	}
	sort.Ints(seats) // we assume the seats are distinct, so all we need is numeric order.

	prev := seats[0]
	for _, seat := range seats[1:] {
		if seat-prev == 2 {
			return prev + 1, nil
		}
		prev = seat
	}
	return 0, fmt.Errorf(`could not locate seat`)
}

func solve1(input string) (max int, err error) {
	paths, err := advent.ReadLines(input)
	if err != nil {
		return
	}
	for _, path := range paths {
		var seat int
		seat, err = locateSeat(path, 128, 8)
		if err != nil {
			return
		}
		if seat > max {
			max = seat
		}
	}
	return
}

func locateSeat(str string, rows, cols int) (seat int, err error) {
	row, col, err := followPath(str, rows, cols)
	if err != nil {
		return
	}
	seat = row*cols + col
	return
}

func followPath(str string, rows, cols int) (row, col int, err error) {
	for _, step := range str {
		switch step {
		case 'B':
			rows /= 2
			row += rows
		case 'F':
			rows /= 2
		case 'R':
			cols /= 2
			col += cols
		case 'L':
			cols /= 2
		default:
			return 0, 0, fmt.Errorf(`cannot follow "%c" in %q`, step, str)
		}
	}
	return
}
