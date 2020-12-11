package main

import (
	"advent"
	"fmt"
)

func main() {
	solution, err := solve2(`input.txt`, 22477624)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve1(file string, n int) (int, error) {
	input, err := advent.ReadInts(file)
	if err != nil {
		return 0, err
	}
	var i, j, k, v int
	for i = n; i < len(input); i++ {
		v = input[i]
		for j = i - n; j < i; j++ {
			for k = j; k < i; k++ {
				if input[j]+input[k] == v {
					goto okay
				}
			}
		}
		return v, nil
	okay:
	}
	return 0, fmt.Errorf(`no anomaly found`)
}

func solve2(file string, x int) (int, error) {
	input, err := advent.ReadInts(file)
	if err != nil {
		return 0, err
	}

	i, j, s, d := 0, 0, 0, 0
	for i = 0; i < len(input); i++ {
		n := input[i]
		min, max := n, n
		for j = i + 1; j < len(input); j++ {
			if n >= x {
				break
			}
			v := input[j]
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			n += v
		}
		if n == x && j-i > d {
			d = j - i
			s = min + max
		}
	}
	return s, nil
}
