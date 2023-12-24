package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	time, distance := parseInput(input)
	input.Close()

	res := calculateWins(time, distance)

	fmt.Printf("Answer: %d\n", res)
}

func calculateWins(time, distance int) int {
	wins := 0

	for speed := 0; speed < time; speed++ {
		timeLeft := time - speed // because speed is equal to time spent waiting

		if speed*timeLeft > distance {
			wins++
		}
	}

	return wins
}

func parseInput(f *os.File) (int, int) {
	fScan := bufio.NewScanner(f)

	fScan.Scan()
	times := findAllNums(fScan.Text())

	fScan.Scan()
	distances := findAllNums(fScan.Text())

	return times, distances
}

func findAllNums(input string) int {
	regex, _ := regexp.Compile("[[:digit:]]+")

	raw := regex.FindAllString(input, -1)
	res, _ := strconv.Atoi(strings.Join(raw, ""))

	return res
}
