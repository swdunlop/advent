package main

import (
	"advent"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	solution, err := solve1(`input.txt`)
	advent.ExitOnErr(err)
	fmt.Println(solution)
}

func solve1(file string) (int, error) {
	return countValidPassports(file, validatorTable{
		`byr`: any, // byr (Birth Year)
		`iyr`: any, // iyr (Issue Year)
		`eyr`: any, // eyr (Expiration Year)
		`hgt`: any, // hgt (Height)
		`hcl`: any, // hcl (Hair Color)
		`ecl`: any, // ecl (Eye Color)
		`pid`: any, // pid (Passport ID)
		`cid`: any, // cid (Country ID)
	})
}

func solve2(file string) (int, error) {
	return countValidPassports(file, validatorTable{
		`byr`: year(1920, 2002), // byr (Birth Year)
		`iyr`: year(2010, 2020), // iyr (Issue Year)
		`eyr`: year(2020, 2030), // eyr (Expiration Year)
		`hgt`: or(
			unit("in", 59, 76),
			unit("cm", 150, 193),
		), // hgt (Height)
		`hcl`: match("#[0-9a-f]{6}"),                // hcl (Hair Color)
		`ecl`: match("amb|blu|brn|gry|grn|hzl|oth"), // ecl (Eye Color)
		`pid`: match("[0-9]{9}"),                    // pid (Passport ID)
		`cid`: any,                                  // cid (Country ID)
	})
}

func or(validators ...func(string) error) func(string) error {
	return func(value string) error {
		for _, validate := range validators {
			err := validate(value)
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf(`is invalid`)
	}
}

func match(pattern string) func(string) error {
	rx := regexp.MustCompile(`^` + pattern + `$`)
	return func(value string) error {
		if !rx.MatchString(value) {
			return fmt.Errorf(`does not match %q`, pattern)
		}
		return nil
	}
}

func year(min, max int) func(string) error {
	return unit("", min, max)
}

func unit(name string, min, max int) func(string) error {
	return func(value string) error {
		if !strings.HasSuffix(value, name) {
			return fmt.Errorf(`missing unit %q`, name)
		}
		value = strings.TrimSuffix(value, name)
		n, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		if n < min || n > max {
			return fmt.Errorf(`%v not in range %v..%v`, n, min, max)
		}
		return nil
	}
}

func any(str string) error {
	if str == "" {
		return fmt.Errorf(`empty value`)
	}
	return nil
}

func countValidPassports(file string, validators validatorTable) (int, error) {
	data, err := advent.ReadBytes(file)
	if err != nil {
		return 0, err
	}
	passports := strings.Split(string(data), "\n\n")
	n := 0
	for _, passport := range passports {
		passport = strings.Replace(passport, "\n", " ", -1)
		_, err := parsePassport(passport, validators)
		if err != nil {
			fmt.Fprintf(os.Stderr, "!! %q is invalid: %v\n", passport, err)
			continue
		}
		fmt.Fprintf(os.Stderr, ".. %q is ok.\n", passport)
		n++
	}
	return n, nil
}

func parsePassport(line string, validators validatorTable) (entry, error) {
	items := strings.Split(line, " ")
	e := make(entry, len(items))
	if len(items) == 0 {
		return nil, fmt.Errorf(`empty line`)
	}
	for _, item := range items {
		if len(item) == 0 {
			continue
		}
		m := strings.SplitN(item, ":", 2)
		if len(m) == 1 {
			return nil, fmt.Errorf(`no colon in %q`, item)
		}
		field, value := m[0], m[1]
		validate, ok := validators[field]
		if !ok {
			return nil, fmt.Errorf(`unsupported field %q in %q`, field, item)
		}
		err := validate(value)
		if err != nil {
			return nil, fmt.Errorf(`%w while validating %q in %q`, err, field, item)
		}
		e[m[0]] = e[m[1]]
	}
	for _, field := range required {
		if _, ok := e[field]; !ok {
			return nil, fmt.Errorf(`missing field %q`, field)
		}
	}
	return e, nil
}

type validatorTable map[string]func(string) error

var required = []string{
	`byr`, `iyr`, `eyr`, `hgt`, `hcl`, `ecl`, `pid`,
	// not: `cid`
}

type entry map[string]string
