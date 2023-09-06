package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
		Algo:
		1. Create a matrix of heights from the input file
	    2. Perform a search in 4 exclusive directions on all the inner cells to determine
	       their visibility
	    3. Count all the visible cells and return
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
	res := parseVisibility(heights)
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

func parseVisibility(heights [][]int) int {
	m, n := len(heights), len(heights[0])

	var checkDirection func(int, int, int, []int) bool
	checkDirection = func(height, i, j int, offset []int) bool {
		if i == 0 || i == m-1 || j == 0 || j == n-1 { // edge cells are always visible
			return true
		}

		new_i, new_j := i+offset[0], j+offset[1]

		if height > heights[new_i][new_j] && checkDirection(height, new_i, new_j, offset) {
			return true
		}

		return false
	}

	res := 2*(m+n) - 4
	offsets := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			for _, offset := range offsets {
				if checkDirection(heights[i][j], i, j, offset) {
					res++
					break
				}
			}
		}
	}

	return res
}

func countVisible(matrix [][]int) int {
	res := 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell == 1 {
				res++
			}
		}
	}

	return res
}

func printMatrix(matrix [][]int) {
	res := ""
	for _, row := range matrix {
		for _, cell := range row {
			res += strconv.Itoa(cell)
		}
		res += "\n"
	}

	fmt.Println(res)
}
