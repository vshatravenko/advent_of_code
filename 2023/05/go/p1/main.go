package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	defaultInputPath = "../input.txt"
	seedsPrefix      = "seeds: "
)

/*
	General idea: build a chain of transformation with each link sorted
	and pass every input element through it

	Algo:
	1. Open the input file and create a set of ranges for each transformer
	2. During filter creation use insertion sort for every range
	3. When the chains are ready, create an array of input values and
	iterate through each chain replacing the prev value with the transformed one
	4. Return the minimum processed number
*/

type Trans struct {
	start int
	end   int
	delta int
}

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

	seeds, transChain := parseInput(input)
	output := processSeeds(transChain, seeds)

	fmt.Printf("Answer: %d\n", minArr(output))
}

func processSeeds(chain [][]Trans, items []int) []int {
	for _, link := range chain {
		for i, item := range items {
			items[i] = processLink(link, item)
		}
	}

	return items
}

func processLink(link []Trans, item int) int {
	for i := 0; i < len(link) && item >= link[i].start; i++ {
		if item < link[i].end {
			return item + link[i].delta
		}
	}

	return item
}

func parseInput(f *os.File) ([]int, [][]Trans) {
	fScan := bufio.NewScanner(f)

	fScan.Scan()
	seeds := parseSeedsLine(fScan.Text())

	transChain := parseChains(fScan)

	return seeds, transChain
}

func parseSeedsLine(line string) []int {
	raw := strings.Split(strings.TrimPrefix(line, seedsPrefix), " ")

	res := make([]int, len(raw))

	for i := 0; i < len(raw); i++ {
		res[i], _ = strconv.Atoi(raw[i])
	}

	return res
}

func parseChains(fScan *bufio.Scanner) [][]Trans {
	chain := [][]Trans{}

	for fScan.Scan() {
		if !strings.Contains(fScan.Text(), "map:") {
			continue
		}

		link := []Trans{}

		fScan.Scan()
		for fScan.Text() != "" {
			trans := parseTransLine(fScan.Text())
			link = insertSorted(link, trans)

			fScan.Scan()
		}

		chain = append(chain, link)
	}

	return chain
}

func parseTransLine(line string) Trans {
	raw := strings.Split(line, " ")
	args := make([]int, 3)
	for i, str := range raw {
		args[i], _ = strconv.Atoi(str)
	}

	return newTrans(args[0], args[1], args[2])
}

func newTrans(dest, source, length int) Trans {
	return Trans{
		start: source,
		end:   source + length,
		delta: dest - source,
	}
}

func insertSorted(arr []Trans, t Trans) []Trans {
	for i := 0; i < len(arr); i++ {
		if t.start < arr[i].start {
			newArr := make([]Trans, len(arr)+1)

			for j := 0; j < i; j++ {
				newArr[j] = arr[j]
			}

			newArr[i] = t

			for j := i; j < len(arr); j++ {
				newArr[j+1] = arr[j]
			}

			return newArr
		}
	}

	return append(arr, t)
}

func minArr(arr []int) int {
	if len(arr) == 0 {
		return math.MinInt
	}

	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}

	return res
}
