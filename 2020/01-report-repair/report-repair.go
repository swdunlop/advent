package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	input := readInputIntegers("2020", "01-report-repair", "input.txt")
	table := make(map[int]struct{}, len(input))
	for _, n := range input {
		table[n] = struct{}{}
	}

	part2(input, table)
}

func part2(input []int, table map[int]struct{}) {
	for i, n := range input {
		for j, m := range input[i:] {
			for _, o := range input[j:] {
				if m+n+o == 2020 {
					fmt.Printf("%v * %v * %v = %v", m, n, o, m*n*o)
					return
				}
			}
		}
	}
}

func part1(input []int, table map[int]struct{}) {
	for i, n := range input {
		for _, m := range input[i:] {
			if m+n == 2020 {
				fmt.Printf("%v * %v = %v", m, n, m*n)
				return
			}
		}
	}
}

func readInputIntegers(path ...string) []int {
	data, err := ioutil.ReadFile(filepath.Join(path...))
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSuffix(data, []byte("\n"))
	lines := strings.Split(string(data), "\n")
	input := make([]int, len(lines))
	for i, line := range lines {
		input[i], err = strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("%w while parsing %q", err, line))
		}
	}
	return input
}
