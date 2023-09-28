package main

/*
	Algo:
	1. Parse the input for each monkey and put it in a struct
	2. Perform the simulation while keeping a counter of ops per monkey
	3. Return the top two counters multiplied
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	defaultInputPath = "input.txt"
	itemDivider      = 3 // used for dividing item after every operation
)

type monkey struct {
	items      []int
	opElems    []string
	divisible  int
	nextMonkey map[bool]int
	opCount    int
}

func main() {
	inputPath := defaultInputPath
	if len(os.Args) == 2 {
		inputPath = os.Args[1]
	}

	f, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("err: could not open %s: %s\n", inputPath, err)
		os.Exit(1)
	}

	monkeys, err := parseInput(f)
	if err != nil {
		fmt.Printf("err: could not parse input: %s\n", err)
		os.Exit(1)
	}
	printMonkeys(monkeys)

	simulate(monkeys)

	res := calculateRes(monkeys)

	fmt.Println("Res:", res)
}

func parseInput(f *os.File) ([]*monkey, error) {
	res := []*monkey{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Monkey") {
			cur, err := parseMonkey(scanner)
			if err != nil {
				return nil, fmt.Errorf("could not parse monkey: %s", err.Error())
			}

			res = append(res, cur)
		}
	}

	return res, nil
}

func parseMonkey(scanner *bufio.Scanner) (*monkey, error) {
	res := &monkey{}

	if ok := scanner.Scan(); !ok || scanner.Text() == "" {
		return nil, fmt.Errorf("monkey input cut too early")
	}
	items, err := parseItemsLine(scanner.Text())
	if err != nil {
		return nil, err
	}
	res.items = items

	if ok := scanner.Scan(); !ok || scanner.Text() == "" {
		return nil, fmt.Errorf("monkey input cut too early")
	}
	res.opElems = strings.Split(strings.TrimPrefix(scanner.Text(), "  Operation: new = "), " ")

	if ok := scanner.Scan(); !ok || scanner.Text() == "" {
		return nil, fmt.Errorf("monkey input cut too early")
	}
	div, err := strconv.Atoi(strings.TrimPrefix(scanner.Text(), "  Test: divisible by "))
	if err != nil {
		return nil, fmt.Errorf("could not convert divisible to num: %s", err)
	}
	res.divisible = div

	if ok := scanner.Scan(); !ok || scanner.Text() == "" {
		return nil, fmt.Errorf("monkey input cut too early")
	}
	trueMonkey, err := strconv.Atoi(strings.TrimPrefix(scanner.Text(), "    If true: throw to monkey "))
	if err != nil {
		return nil, fmt.Errorf("could not parse next monkey num: %s", err)
	}

	if ok := scanner.Scan(); !ok || scanner.Text() == "" {
		return nil, fmt.Errorf("monkey input cut too early")
	}
	falseMonkey, err := strconv.Atoi(strings.TrimPrefix(scanner.Text(), "    If false: throw to monkey "))
	if err != nil {
		return nil, fmt.Errorf("could not parse next monkey num: %s", err)
	}
	res.nextMonkey = map[bool]int{true: trueMonkey, false: falseMonkey}

	return res, nil
}

func parseItemsLine(line string) ([]int, error) {
	itemsRaw := strings.Split(strings.TrimPrefix(line, "  Starting items: "), ", ")
	items := make([]int, len(itemsRaw))
	for i, item := range itemsRaw {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, fmt.Errorf("could not convert starting item %s to num: %s", item, err)
		}

		items[i] = num
	}

	return items, nil
}

func simulate(monkeys []*monkey) error {
	for i := 0; i < 20; i++ {
		for idx, m := range monkeys {
			for len(m.items) > 0 {
				prevItem := m.items[0]
				newItem, err := m.performOp()
				if err != nil {
					return err
				}
				testRes := newItem%m.divisible == 0
				newIdx := m.nextMonkey[testRes]
				monkeys[newIdx].appendItem(newItem)
				fmt.Printf("monkey %d: prev item: %d; new item: %d; new monkey: %d\n", idx, prevItem, newItem, newIdx)
			}
		}

		fmt.Println()
		printMonkeys(monkeys)
	}

	return nil
}

func (m *monkey) performOp() (int, error) {
	if len(m.items) == 0 {
		return -1, fmt.Errorf("no items left")
	}

	res := m.items[0]
	if len(m.items) == 1 {
		m.items = []int{}
	} else {
		m.items = m.items[1:]
	}

	var first, second int
	var err error

	if m.opElems[0] == "old" {
		first = res
	} else {
		first, err = strconv.Atoi(m.opElems[0])
		if err != nil {
			return -1, fmt.Errorf("could not convert %s to num: %s", m.opElems[0], err)
		}
	}

	if m.opElems[2] == "old" {
		second = res
	} else {
		second, err = strconv.Atoi(m.opElems[2])
		if err != nil {
			return -1, fmt.Errorf("could not convert %s to num: %s", m.opElems[2], err)
		}
	}

	switch m.opElems[1] {
	case "+":
		res = first + second
	case "-":
		res = first - second
	case "*":
		res = first * second
	case "/":
		res = first / second
	default:
		return -1, fmt.Errorf("invalid operator: %s", m.opElems[1])
	}

	m.opCount++

	return res / 3, nil
}

func (m *monkey) appendItem(item int) {
	m.items = append(m.items, item)
}

func calculateRes(monkeys []*monkey) int {
	firstMax, secondMax := math.MinInt, math.MinInt

	for _, m := range monkeys {
		if m.opCount > firstMax {
			if firstMax > secondMax {
				secondMax = firstMax
			}
			firstMax = m.opCount
		} else if m.opCount > secondMax {
			secondMax = m.opCount
		}
	}

	fmt.Printf("top opCounts: %d + %d\n", firstMax, secondMax)

	return firstMax * secondMax
}

func printMonkeys(monkeys []*monkey) {
	for i, m := range monkeys {
		fmt.Printf("monkey %d: %+v\n", i, m)
	}
}
