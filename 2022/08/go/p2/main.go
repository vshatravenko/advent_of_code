package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
		Algo:
		1. Create a matrix of heights from the input file
	    2. Count the visible distance on every direction for each tree, keeping a maximum
	    3. Return the result
*/

const (
	inputFile = "input.txt"
)

func main() {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Could not open the input file %s: %s\n", inputFile, err)
	}

	heights := parseHeights(input)
	res := getMaxScore(heights)
	fmt.Println("Result:", res)
}

func parseHeights(f *os.File) [][]int {
	res := [][]int{}
	fScanner := bufio.NewScanner(f)
	i := 0

	for fScanner.Scan() {
		line := fScanner.Text()
		n := len(line)

		res = append(res, make([]int, n))
		for j := 0; j < n; j++ {
			res[i][j], _ = strconv.Atoi(string(line[j]))
		}

		i++
	}

	return res
}

func getMaxScore(heights [][]int) int {
	m, n := len(heights), len(heights[0])

	var countDepth func(int, int, int, []int) int
	countDepth = func(height, i, j int, offset []int) int {
		if i == 0 || i == m-1 || j == 0 || j == n-1 {
			return 0
		}

		new_i, new_j := i+offset[0], j+offset[1]

		if height > heights[new_i][new_j] {
			return 1 + countDepth(height, new_i, new_j, offset)
		}

		return 1
	}

	maxScore := math.MinInt
	offsets := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			depths := make([]int, len(offsets))
			for k, offset := range offsets {
				depths[k] = countDepth(heights[i][j], i, j, offset)
			}

			score := 1
			for _, d := range depths {
				score *= d
			}
			if score > maxScore {
				maxScore = score
			}

		}
	}

	return maxScore
}
