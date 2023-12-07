package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			if line[i] >= '0' && line[i] <= '9' {
				digit, _ := strconv.Atoi(line[i : i+1])
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

		res += first*10 + last
	}

	return res
}
