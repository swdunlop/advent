package main

import (
	"advent"
	"fmt"
	"strings"
)

func main() {
	solution, err := solve2(`input.txt`)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve1(file string) (int, error) {
	groups, err := readGroups(file)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, g := range groups {
		n += g.anyTrue().trueCount()
	}
	return n, nil
}

func solve2(file string) (int, error) {
	groups, err := readGroups(file)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, g := range groups {
		n += g.allTrue().trueCount()
	}
	return n, nil
}
func readGroups(file string) ([]group, error) {
	data, err := advent.ReadBytes(file)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	groups := make([]group, len(parts))
	for i, str := range parts {
		groups[i], err = parseGroup(str)
		if err != nil {
			return nil, err
		}
	}
	return groups, nil
}

func parseGroup(str string) (g group, err error) {
	lines := strings.Split(str, "\n")
	g = make(group, len(lines))
	for i, line := range lines {
		g[i], err = parseForm(line)
		if err != nil {
			return nil, fmt.Errorf(`%w while parsing %q`, err, line)
		}
	}
	return g, nil
}

type group []*form

func (g group) anyTrue() *form {
	table := new(form)
	for i := 0; i < 26; i++ {
		for _, form := range g {
			if form[i] {
				table[i] = true
				break
			}
		}
	}
	return table
}

func (g group) allTrue() *form {
	table := new(form)
	for i := 0; i < 26; i++ {
		for _, form := range g {
			table[i] = true
			if !form[i] {
				table[i] = false
				break
			}
		}
	}
	return table
}
func parseForm(str string) (*form, error) {
	f := new(form)
	for _, ch := range str {
		if ch < 'a' || ch > 'z' {
			return nil, fmt.Errorf(`"%c" is not a question on the form`, ch)
		}
		f[ch-'a'] = true
	}
	return f, nil
}

type form [26]bool

func (f *form) trueCount() int {
	n := 0
	for _, a := range *f {
		if a {
			n++
		}
	}
	return n
}
