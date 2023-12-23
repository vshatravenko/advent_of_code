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

	P2 changes: work with seed ranges, apply transformations to the whole range
	and split the input range when it's not applicable
	e.g. 52:10 range applied to a 53:52:5 filter would be split into 53:5 & 57:5
	In the end of the transformation chain, the answer is the smallest start of all the ranges

	Algo:
	1. Open the input file and create a set of ranges for each transformer
	2. During filter creation use insertion sort for every range
	3. When the chains are ready, create an array of input values and
	iterate through each chain replacing the prev value with the transformed one
	4. Return the minimum processed number
*/

type Range struct {
	start  int
	length int
}

type Trans struct {
	Range *Range
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

func processSeeds(chain [][]*Trans, items []*Range) []*Range {
	for _, link := range chain {
		buffer := []*Range{}
		for _, item := range items {
			buffer = append(buffer, processLink(link, item)...)
		}

		items = buffer
	}

	return items
}

func processLink(link []*Trans, item *Range) []*Range {
	buffer := []*Range{item}
	res := []*Range{}
	for i := 0; i < len(buffer); i++ {
		for j := 0; j < len(link); j++ {
			if isOverlap(buffer[i], link[j].Range) {
				pre, overlap, post := extractOverlap(buffer[i], link[j].Range, link[j].delta)

				res = append(res, overlap)

				if pre != nil {
					buffer = append(buffer, pre)
				}

				if post != nil {
					buffer = append(buffer, post)
				}

				break
			}

			if j == len(link)-1 {
				res = append(res, buffer[i])
			}
		}
	}

	return res
}

func isOverlap(a, b *Range) bool {
	return min(a.start+a.length-1, b.start+b.length-1)-max(a.start, b.start) > 0
}

func extractOverlap(a, b *Range, delta int) (*Range, *Range, *Range) {
	if a.start == b.start && a.length == b.length {
		return nil, a, nil
	}

	var pre, overlap, post *Range = nil, nil, nil

	if b.start-a.start > 0 {
		pre = &Range{start: a.start, length: b.start - a.start}
	}

	aEnd, bEnd := a.start+a.length-1, b.start+b.length-1
	overlap = &Range{start: max(a.start, b.start) + delta, length: min(aEnd, bEnd) - max(a.start, b.start) + 1}

	if aEnd-bEnd > 0 {
		post = &Range{start: bEnd, length: aEnd - bEnd}
	}

	return pre, overlap, post
}

func parseInput(f *os.File) ([]*Range, [][]*Trans) {
	fScan := bufio.NewScanner(f)

	fScan.Scan()
	seeds := parseSeedsLine(fScan.Text())

	transChain := parseChains(fScan)

	return seeds, transChain
}

func parseSeedsLine(line string) []*Range {
	raw := strings.Split(strings.TrimPrefix(line, seedsPrefix), " ")

	res := []*Range{}

	for i := 0; i < len(raw)-1; i += 2 {
		start, _ := strconv.Atoi(raw[i])
		length, _ := strconv.Atoi(raw[i+1])
		res = append(res, &Range{start: start, length: length})
	}

	return res
}

func parseChains(fScan *bufio.Scanner) [][]*Trans {
	chain := [][]*Trans{}

	for fScan.Scan() {
		if !strings.Contains(fScan.Text(), "map:") {
			continue
		}

		link := []*Trans{}

		fScan.Scan()
		for fScan.Text() != "" {
			trans := parseTransLine(fScan.Text())
			link = append(link, trans)

			fScan.Scan()
		}

		chain = append(chain, link)
	}

	return chain
}

func parseTransLine(line string) *Trans {
	raw := strings.Split(line, " ")
	args := make([]int, 3)
	for i, str := range raw {
		args[i], _ = strconv.Atoi(str)
	}

	return newTrans(args[0], args[1], args[2])
}

func newTrans(dest, start, length int) *Trans {
	return &Trans{
		Range: &Range{
			start:  start,
			length: length,
		},
		delta: dest - start,
	}
}

func minArr(arr []*Range) int {
	if len(arr) == 0 {
		return math.MinInt
	}

	res := arr[0].start
	for i := 1; i < len(arr); i++ {
		if arr[i].start < res {
			res = arr[i].start
		}
	}

	return res
}
