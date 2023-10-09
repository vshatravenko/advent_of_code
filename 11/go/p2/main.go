package main

/*
	Algo:
	1. Parse the input for each monkey and put it in a struct
	2. Perform the simulation while keeping a counter of ops per monkey
	3. Return the top two counters multiplied

	The main change now is that the simulation spans 10000 rounds
    and the worry levels per item are not divided by 3, so int would overflow
	math/big.Int is used instead
*/

/*
	Thoughts
	Currently, even with math/big in play, the simulation stalls at the first ~10%
	just because of the fact how huge the numbers are, with operations
	starting to take multiple seconds.
	All these computations are done for just one thing - determining whether an item is
	divisible by a given number and than transforming it by adding or multiplying it
	I suppose that we could circumvent the operational complexity of big numbers by
	either checking the divisibility using a set of operations on a number's integers,
	or by keeping a set of factors (e.g. (((5 * 6) + 10) * 6) + 11) etc. and then unwrapping it
	but it sounds a bit tough
	Also, I may be approaching the problem from the wrong direction, and divisibility
	shouldn't even come into play. Maybe, there's a more efficient way of determining
	the next monkey.

	A possible clue - all the divisibility numbers are prime

	Maybe, we could keep the list of operands and check their individual adherence to divisibility?
*/

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	defaultInputPath = "input.txt"
	simulationLen    = 10000
)

type monkey struct {
	items      []*big.Int
	opElems    []string
	divisor    *big.Int
	nextMonkey map[bool]int
	opCount    *big.Int
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
	res := &monkey{opCount: big.NewInt(0)}

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
	div, err := strconv.ParseInt(strings.TrimPrefix(scanner.Text(), "  Test: divisible by "), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not convert divisible to num: %s", err)
	}
	res.divisor = big.NewInt(div)

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

func parseItemsLine(line string) ([]*big.Int, error) {
	itemsRaw := strings.Split(strings.TrimPrefix(line, "  Starting items: "), ", ")
	items := make([]*big.Int, len(itemsRaw))
	for i, item := range itemsRaw {
		num, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("could not convert starting item %s to num: %s", item, err)
		}

		items[i] = big.NewInt(num)
	}

	return items, nil
}

func simulate(monkeys []*monkey) error {
	lcm := calcLCM(monkeys)

	for i := 0; i < simulationLen; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				newItem, err := m.performOp(lcm)
				if err != nil {
					return err
				}

				modRes := &big.Int{}
				testRes := modRes.Mod(newItem, m.divisor).Cmp(big.NewInt(0)) == 0
				newIdx := m.nextMonkey[testRes]
				monkeys[newIdx].appendItem(newItem)
			}
		}

		fmt.Println()
		printMonkeys(monkeys)
	}

	return nil
}

func (m *monkey) performOp(threshold *big.Int) (*big.Int, error) {
	if len(m.items) == 0 {
		return &big.Int{}, fmt.Errorf("no items left")
	}

	old := m.items[0]
	if len(m.items) == 1 {
		m.items = []*big.Int{}
	} else {
		m.items = m.items[1:]
	}

	first, second := &big.Int{}, &big.Int{}

	if m.opElems[0] == "old" {
		first = old
	} else {
		raw, err := strconv.ParseInt(m.opElems[0], 10, 64)
		if err != nil {
			return &big.Int{}, fmt.Errorf("could not convert %s to num: %s", m.opElems[0], err)
		}

		first = big.NewInt(raw)
	}

	if m.opElems[2] == "old" {
		second = old
	} else {
		raw, err := strconv.ParseInt(m.opElems[2], 10, 64)
		if err != nil {
			return &big.Int{}, fmt.Errorf("could not convert %s to num: %s", m.opElems[2], err)
		}

		second = big.NewInt(raw)
	}

	res := &big.Int{}
	switch m.opElems[1] {
	case "+":
		res.Add(first, second) // pretty unusual, the result is stored in the caller(res)
	case "-":
		res.Sub(first, second)
	case "*":
		res.Mul(first, second)
	case "/":
		res.Div(first, second)
	default:
		return &big.Int{}, fmt.Errorf("invalid operator: %s", m.opElems[1])
	}

	m.opCount.Add(m.opCount, big.NewInt(1))
	res = res.Mod(res, threshold)

	return res, nil
}

func (m *monkey) appendItem(item *big.Int) {
	m.items = append(m.items, item)
}

// LCM means lowest common multiplier, a number which all the participating numbers divide into
// Since all the divisibility numbers are primes, we simply need to multiply them between each other
func calcLCM(monkeys []*monkey) *big.Int {
	res := big.NewInt(1)

	for _, monkey := range monkeys {
		res = res.Mul(res, monkey.divisor)
	}

	return res
}

func calculateRes(monkeys []*monkey) *big.Int {
	firstMax, secondMax := big.NewInt(0), big.NewInt(0)

	for _, m := range monkeys {
		if m.opCount.Cmp(firstMax) == 1 { // Cmp() => 1 - greater than, 0 - equal, -1 - less than
			if firstMax.Cmp(secondMax) == 1 {
				secondMax = firstMax
			}
			firstMax = m.opCount
		} else if m.opCount.Cmp(secondMax) == 1 {
			secondMax = m.opCount
		}
	}

	fmt.Printf("top opCounts: %d & %d\n", firstMax, secondMax)

	return firstMax.Mul(firstMax, secondMax)
}

func printMonkeys(monkeys []*monkey) {
	for i, m := range monkeys {
		fmt.Printf("monkey %d: opCount: %v\n", i, *m.opCount)
	}
}
