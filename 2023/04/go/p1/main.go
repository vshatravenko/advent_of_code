package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultInputPath = "../input.txt"
)

func main() {
	inputPath := defaultInputPath
	if len(os.Args) == 2 {
		inputPath = os.Args[1]
	}

	input, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("err: could not open %s: %s", inputPath, err.Error())
		os.Exit(1)
	}

	fmt.Printf("Answer: %d\n", solve(input))
}

func solve(f *os.File) int {
	fScanner := bufio.NewScanner(f)

	res := 0

	for fScanner.Scan() {
		line := fScanner.Text()
		sides := parseCard(line)

		winningNums := createSet(sides[1])

		matchCount := 0
		for _, num := range sides[0] {
			if _, ok := winningNums[num]; ok {
				matchCount++
			}
		}

		if matchCount > 0 {
			res += 1 << (matchCount - 1)
		}
	}

	return res
}

func parseCard(line string) [][]int {
	colonLocation := strings.Index(line, ":") // since ":" is where the card prefix ends

	sidesRaw := strings.Split(line[colonLocation:], "|")
	res := [][]int{}

	for _, sideRaw := range sidesRaw {
		side := []int{}
		for _, elem := range strings.Split(sideRaw, " ") {
			if elem != "" {
				elemParsed, _ := strconv.Atoi(elem)
				side = append(side, elemParsed)
			}
		}

		res = append(res, side)
	}

	return res
}

func createSet(input []int) map[int]bool {
	set := make(map[int]bool)

	for _, n := range input {
		set[n] = true
	}

	return set
}
