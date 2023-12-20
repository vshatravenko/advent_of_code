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

	lines := readAllLines(input)

	fmt.Printf("Answer: %d\n", solve(lines))
}

func solve(lines []string) int {
	cardCount := len(lines)
	deck := make([]int, cardCount)
	for i := 0; i < cardCount; i++ {
		deck[i] = 1
	}

	for i, line := range lines {
		sides := parseCard(line)
		winningNums := createSet(sides[1])
		matchCount := 0

		for _, num := range sides[0] {
			if _, ok := winningNums[num]; ok {
				matchCount++
			}
		}

		for j := 1; j <= matchCount; j++ {
			deck[i+j] += 1 * deck[i]
		}
	}

	return countDeck(deck)
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

func countDeck(deck []int) int {
	res := 0
	for _, count := range deck {
		res += count
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

func readAllLines(f *os.File) []string {
	fScanner := bufio.NewScanner(f)

	lines := []string{}
	for fScanner.Scan() {
		lines = append(lines, fScanner.Text())
	}

	return lines
}
