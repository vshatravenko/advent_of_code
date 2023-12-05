package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultInputPath = "input.txt"
)

type Op struct {
	name   string
	value  int
	cycles int
}

func main() {
	inputPath := defaultInputPath
	if len(os.Args) == 2 {
		inputPath = os.Args[1]
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Could not open %s: %s\n", inputPath, err)
		os.Exit(1)
	}

	// parse input commands
	ops := parseInput(inputFile)

	// run the simulation
	res := simulate(ops)

	fmt.Printf("Result: %d\n", res)
}

func parseInput(f *os.File) []*Op {
	res := []*Op{}
	fScanner := bufio.NewScanner(f)

	for fScanner.Scan() {
		tokens := strings.Split(fScanner.Text(), " ")

		if len(tokens) == 1 {
			res = append(res, &Op{name: "noop", cycles: 1})
		} else {
			value, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Printf("err: could not parse line %s: %s\n", fScanner.Text(), err)
				os.Exit(1)
			}

			res = append(res, &Op{name: tokens[0], value: value, cycles: 2})
		}
	}

	return res
}

/*
Would run the simulation and return the sum of
the 20th, 60th, 100th, 140th, 180th, and 220th
frequency strengths(register val * cycle)

FIXME: fugly
*/
func simulate(ops []*Op) int {
	opsLen := len(ops)
	sum := 0
	registerVal := 1
	opIdx := 0
	pendingVal := 0
	isPendingAddx := false

	for cycle := 1; cycle <= 220 && opIdx < opsLen; cycle++ {
		ops[opIdx].cycles--

		if isPendingAddx {
			registerVal += pendingVal
			isPendingAddx = false
		}
		fmt.Printf("cycle: %d; register: %d; op: %+v\n", cycle, registerVal, ops[opIdx])

		if ops[opIdx].cycles == 0 {
			if ops[opIdx].name == "addx" {
				pendingVal = ops[opIdx].value
				isPendingAddx = true
			}

			opIdx++
		}

		if cycle == 20 || (cycle-20)%40 == 0 {
			sum += registerVal * cycle
			fmt.Printf("Added %d to sum, result: %d\n", registerVal*cycle, sum)
		}
	}

	return sum
}
