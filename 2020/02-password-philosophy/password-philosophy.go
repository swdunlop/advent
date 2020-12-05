package main

import (
	"advent"
	"fmt"
	"unicode/utf8"
)

func main() {
	solution, err := solve2("input.txt")
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve1(file string) (int, error) {
	n := 0
	err := forEachEntry(file, func(e *entry) error {
		o := 0
		for _, ch := range e.Password {
			if ch == e.Ch {
				o++
			}
		}
		if o >= e.A && o <= e.B {
			n++
		}
		return nil
	})
	return n, err
}

func solve2(file string) (int, error) {
	n := 0
	err := forEachEntry(file, func(e *entry) error {
		if e.checkCh(e.A-1) != e.checkCh(e.B-1) {
			n++
		}
		return nil
	})
	return n, err
}

func (e *entry) checkCh(ofs int) bool {
	if ofs < 0 || ofs >= len(e.Password) {
		return false
	}
	r, _ := utf8.DecodeRuneInString(e.Password[ofs:])
	return r == e.Ch
}

func forEachEntry(file string, do func(e *entry) error) error {
	var e entry
	lines, err := advent.ReadLines(file)
	if err != nil {
		return err
	}
	for _, line := range lines {
		n, err := fmt.Sscanf(line, "%v-%v %c: %v", &e.A, &e.B, &e.Ch, &e.Password)
		if err != nil {
			return fmt.Errorf(`%w in field %v while parsing %q`, err, n+1, line)
		}
		err = do(&e)
		if err != nil {
			return fmt.Errorf(`%w while processing %q`, err, line)
		}
	}
	return nil
}

type entry struct {
	A, B     int
	Ch       rune
	Password string
}
