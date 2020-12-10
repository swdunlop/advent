package main

import (
	"advent"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	solution, err := solve2(`input.txt`)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve1(file string) (int, error) {
	rs, err := readRuleset(file)
	if err != nil {
		return 0, err
	}
	cs := make(colorset, len(rs))
	rs.identifyOutermostColors(rs[`shiny gold`], cs)
	return len(cs), nil
}

func solve2(file string) (int, error) {
	rs, err := readRuleset(file)
	if err != nil {
		return 0, err
	}
	return rs.countContents(`shiny gold`) - 1, nil // you already own the gold bag!
}

func readRuleset(file string) (ruleset, error) {
	lines, err := advent.ReadLines(file)
	if err != nil {
		return nil, err
	}
	rs := make(ruleset, len(lines))
	for _, line := range lines {
		r, err := parseRule(line)
		if err != nil {
			return nil, fmt.Errorf(`%w in %q`, err, line)
		}
		if r != nil {
			rs[r.color] = r
		}
	}
	for _, r := range rs {
		for color := range r.contents {
			r2, ok := rs[color]
			if !ok {
				return nil, fmt.Errorf(`missing rule for %q`, color)
			}
			r2.containers = append(r2.containers, r)
		}
	}
	return rs, nil
}

type ruleset map[string]*rule

func (rs ruleset) countContents(color string) int {
	r := rs[color]
	if r == nil {
		return 1
	}
	if r.counted {
		return r.count
	}
	n := 1 // remember to count your own donkey, Johua.
	for color, count := range r.contents {
		n += count * rs.countContents(color)
	}
	r.counted, r.count = true, n
	return n
}

func (rs ruleset) identifyOutermostColors(r *rule, cs colorset) {
	if cs.has(r.color) {
		return
	}
	for _, r2 := range r.containers {
		rs.identifyOutermostColors(r2, cs)
		cs.add(r2.color)
	}
}

type colorset map[string]struct{}

func (cs colorset) has(str string) bool { _, ok := cs[str]; return ok }
func (cs colorset) add(str string)      { cs[str] = struct{}{} }

func parseRule(str string) (*rule, error) {
	m := strings.SplitN(str, ` bags contain `, 2)
	if len(m) != 2 {
		return nil, fmt.Errorf(`missing " bags contain "`)
	}

	items := strings.Split(strings.TrimSuffix(m[1], "."), ", ")
	r := &rule{
		color:      m[0],
		containers: make([]*rule, 0, 8),
		contents:   make(map[string]int, len(items)),
	}
	if m[1] == "no other bags." {
		return r, nil
	}

	for _, item := range items {
		m := strings.SplitN(item, " ", 4)
		if len(m) != 4 {
			return nil, fmt.Errorf(`%q has less than four words`, item)
		}
		n, err := strconv.Atoi(m[0])
		if err != nil {
			if m[0] == "one" {
				n = 1
			} else {
				return nil, fmt.Errorf(`%w while parsing %q`, err, item)
			}
		}
		r.contents[m[1]+" "+m[2]] = n
	}
	return r, nil
}

// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// ^--color---^              ^----item-------^  ^------item-------^
//                     count-^ ^--color-^

type rule struct {
	color      string
	containers []*rule
	contents   map[string]int
	count      int
	counted    bool
}
