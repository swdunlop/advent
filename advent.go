package advent

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ExitOnErr exits with status code one after printing an error, if err is non-nil.
func ExitOnErr(err error) {
	if err != nil {
		println(`!!`, err.Error())
		os.Exit(1)
	}
}

// ReadInts uses ReadLines to read a slice of base 10 integers.
func ReadInts(path ...string) ([]int, error) {
	lines, err := ReadLines(path...)
	if err != nil {
		return nil, err
	}
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf(`%w while parsing %q in line %v`, err, line, i+1)
		}
	}
	return ints, nil
}

// ReadLines uses ReadBytes to read a slice of lines which must be terminated by "\n".
func ReadLines(path ...string) ([]string, error) {
	data, err := ReadBytes(path...)
	if err != nil {
		return nil, err
	}
	str := string(data)
	str = strings.TrimSuffix(str, "\n")
	return strings.Split(str, "\n"), nil
}

// ReadBytes reads the file from the provided path joined using filepath.
func ReadBytes(path ...string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(path...))
}
