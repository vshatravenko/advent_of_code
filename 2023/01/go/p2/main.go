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

	fmt.Printf("Answer: %d\n", processInput(input))
}

func processInput(f *os.File) int {
	res := 0
	fScanner := bufio.NewScanner(f)

	for fScanner.Scan() {
		line := fScanner.Text()
		first, last := -1, -1

		for i := 0; i < len(line); i++ {
			digit := -1
			if line[i] >= '1' && line[i] <= '9' {
				digit, _ = strconv.Atoi(line[i : i+1])
			} else if parsed, _, ok := parseDigitWord(line[i:]); ok {
				digit = parsed
			}

			if digit != -1 {
				if first == -1 {
					first = digit
				} else {
					last = digit
				}
			}
		}

		if last == -1 {
			last = first
		}

		fmt.Printf("Parsing line: %s, Found %d + %d\n", line, first, last)

		res += first*10 + last
	}

	return res
}

func parseDigitWord(input string) (int, int, bool) {
	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for word, digit := range words {
		if strings.HasPrefix(input, word) {
			return digit, len(word), true
		}
	}

	return 0, 0, false
}
