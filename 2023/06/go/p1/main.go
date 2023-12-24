package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	times, distances := parseInput(input)
	input.Close()

	res := simulateRaces(times, distances)

	fmt.Printf("Answer: %d\n", multiplyArr(res))
}

func simulateRaces(times, distances []int) []int {
	raceCount := len(times)
	winWays := make([]int, raceCount)

	for i := 0; i < raceCount; i++ {
		winWays[i] = calculateWins(times[i], distances[i])
	}

	return winWays
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

func parseInput(f *os.File) ([]int, []int) {
	fScan := bufio.NewScanner(f)

	fScan.Scan()
	times := findAllNums(fScan.Text())

	fScan.Scan()
	distances := findAllNums(fScan.Text())

	return times, distances
}

func findAllNums(input string) []int {
	regex, _ := regexp.Compile("[[:digit:]]+")

	raw := regex.FindAllString(input, -1)

	res := make([]int, len(raw))
	for i, str := range raw {
		res[i], _ = strconv.Atoi(str)
	}

	return res
}

func multiplyArr(arr []int) int {
	res := 1
	for _, elem := range arr {
		res *= elem
	}

	return res
}
