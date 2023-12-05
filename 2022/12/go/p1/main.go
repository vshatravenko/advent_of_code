package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
	The goal for this day is to find the minimum path
	from the starting point to the goal

	One way to accomplish that would be to use DFS
	and track the distance for every path, selecting the smallest position

	Otherwise, it could be solved by using DP to build a matrix
	of positions with every cell denoting successful path length

	Algo:
	1.  Parse the file as a 2D matrix
	2.  Start the DFS process in four directions
	    For every new cell in the path, increment the distance
		If there is a valid path, return its length
		If not, return -1
	3. Return the minimum path from start to finish

*/

const (
	defaultInputPath = "input.txt"
	startChar        = 'S'
	endChar          = 'E'
)

var (
	endRow int
	endCol int
)

func main() {
	inputPath := defaultInputPath
	if len(os.Args) == 2 {
		inputPath = os.Args[1]
	}
	fmt.Println("Parsing", inputPath)

	f, err := os.Open(inputPath)
	if err != nil {
	}

	matrix := parseInput(f)
	f.Close()

	startRow, startCol := findStartPos(matrix)
	matrix[startRow][startCol] = 'a'

	endRow, endCol = findEndPos(matrix)
	matrix[endRow][endCol] = 'z'

	res := findMinPathLen(startRow, startCol, 0, matrix)

	fmt.Println("result:", res)
}

func parseInput(f *os.File) [][]byte {
	res := [][]byte{}
	fScanner := bufio.NewScanner(f)

	for fScanner.Scan() {
		line := fScanner.Bytes()
		res = append(res, line)
	}

	return res
}

func findStartPos(matrix [][]byte) (int, int) {
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == startChar {
				return i, j
			}
		}
	}

	return -1, -1
}

func findEndPos(matrix [][]byte) (int, int) {
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == endChar {
				return i, j
			}
		}
	}

	return -1, -1
}

func findMinPathLen(row, col, curLen int, matrix [][]byte) int {
	if row == endRow && col == endCol {
		return curLen
	}

	printMatrix(matrix)
	curChar := matrix[row][col]
	matrix[row][col] = '#'

	res := math.MaxInt
	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, offset := range offsets {
		newRow, newCol := row+offset[0], col+offset[1]

		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) || matrix[newRow][newCol] == '#' {
			continue
		}

		dist := abs(int(curChar) - int(matrix[newRow][newCol]))
		if dist == 0 || dist == 1 {
			pathLen := findMinPathLen(newRow, newCol, curLen+1, matrix)

			if pathLen == -1 {
				continue
			}

			res = min(res, pathLen)
		}
	}

	matrix[row][col] = curChar

	if res != math.MaxInt {
		return res
	}

	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func printMatrix(matrix [][]byte) {
	for _, row := range matrix {
		str := ""
		for _, c := range row {
			str += string(c)
		}
		fmt.Println(str)
	}
}
