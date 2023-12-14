package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	Algo:
	1. Convert file output into a 2D byte arr
	2. Extract number start and end in a given row
	3. Check each byte in array for surrounding engine parts
	4. If there are any, get a byte slice a convert it to int, add to total sum
*/

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

	raw, err := io.ReadAll(input)
	if err != nil {
		fmt.Printf("err: could not read %s file contents: %s", inputPath, err.Error())
		os.Exit(1)
	}

	if raw[len(raw)-1] == '\n' {
		raw = raw[:len(raw)-1]
	}

	fmt.Printf("Answer: %d\n", extractEngineParts(string(raw)))
}

func extractEngineParts(raw string) int {
	matrix := strings.Split(raw, "\n")

	fmt.Printf("parsed matrix: %v\n", matrix)

	rows, cols := len(matrix), len(matrix[0])
	fmt.Printf("rows: %d; cols: %d\n", rows, cols)

	sum := 0
	for i := 0; i < rows; i++ {
		numStart, numEnd := -1, -1

		for j := 0; j < cols; j++ {
			if isDigit(matrix[i][j]) {
				numStart = j

				for j < cols && isDigit(matrix[i][j]) {
					j++
				}

				numEnd = j - 1

				if checkEnginePart(matrix, i, numStart, numEnd) {
					num, _ := strconv.Atoi(matrix[i][numStart : numEnd+1])
					// fmt.Printf("detected part: %d\n", num)
					sum += num
				}

				numStart, numEnd = -1, -1
			}
		}
	}

	return sum
}

func checkEnginePart(matrix []string, row, start, end int) bool {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

	rows, cols := len(matrix), len(matrix[0])

	for col := start; col <= end; col++ {
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				if !isDigit(matrix[newRow][newCol]) && matrix[newRow][newCol] != '.' {
					return true
				}
			}
		}
	}

	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
