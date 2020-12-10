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
	var vm vm
	var err error
	vm.program, err = readProgram(file)
	if err != nil {
		return 0, err
	}
	for {
		if vm.program[vm.pc].visited {
			break
		}
		err = vm.step()
		if err != nil {
			return 0, err
		}
	}
	return vm.acc, nil
}

func solve2(file string) (int, error) {
	var vm vm
	var err error
	vm.program, err = readProgram(file)
	if err != nil {
		return 0, err
	}
	for pp := 0; pp < len(vm.program); pp++ {
		op := vm.program[pp].operation
		if op == jmp {
			vm.program[pp].operation = nop
		} else if op == nop {
			vm.program[pp].operation = jmp
		} else {
			continue
		}

		vm.pc, vm.acc = 0, 0
		for i := range vm.program {
			vm.program[i].visited = false
		}

		for {
			if vm.pc >= len(vm.program) {
				return vm.acc, nil
			}
			in := vm.program[vm.pc]
			if in.visited {
				break
			}
			err = vm.step()
			if err != nil {
				return 0, err
			}
		}

		vm.program[pp].operation = op
	}
	return 0, fmt.Errorf(`no patch found`)
}

type vm struct {
	pc      int
	acc     int
	program []instruction
}

func (vm *vm) step() error {
	inst := &vm.program[vm.pc]
	inst.visited = true
	switch inst.operation {
	case acc:
		vm.acc += inst.argument
		vm.pc++
	case jmp:
		if inst.argument == 0 {
			return fmt.Errorf(`infinite loop detected at %v`, vm.pc)
		}
		vm.pc += inst.argument
	case nop:
		vm.pc++
	default:
		return fmt.Errorf(`unsupported instruction %v at %v`, inst.operation, vm.pc)
	}
	return nil
}

func readProgram(file string) ([]instruction, error) {
	lines, err := advent.ReadLines(file)
	program := make([]instruction, len(lines))
	for pc, line := range lines {
		program[pc], err = parseInstruction(line)
		if err != nil {
			return nil, fmt.Errorf(`%w in %q`, err, line)
		}
	}
	return program, nil
}

func parseInstruction(str string) (in instruction, err error) {
	m := strings.SplitN(str, " ", 3)
	if len(m) != 2 {
		return in, fmt.Errorf(`malformed instruction`)
	}
	op, ok := operations[m[0]]
	if !ok {
		return in, fmt.Errorf(`%q is not an opcode`, m[0])
	}
	in.operation = op
	in.argument, err = strconv.Atoi(m[1])
	return in, err
}

type instruction struct {
	operation operation
	argument  int
	visited   bool
}

type operation int

var operations = map[string]operation{
	`nop`: nop,
	`acc`: acc,
	`jmp`: jmp,
}

const (
	nop = operation(iota)
	acc
	jmp
)
