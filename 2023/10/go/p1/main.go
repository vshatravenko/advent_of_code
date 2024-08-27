package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	defaultInputPath = "../input.txt"
	startChar        = 'S'
	markChar         = '#'
)

var (
	// {from, to}
	pipeMap = map[rune][]rune{
		'|': {'S', 'N'},
		'-': {'W', 'E'},
		'L': {'E', 'N'},
		'J': {'W', 'N'},
		'7': {'W', 'S'},
		'F': {'S', 'E'},
	}

	directions = []rune{'N', 'S', 'W', 'E'}

	dirMap = map[rune][]int{
		'N': {-1, 0},
		'S': {1, 0},
		'W': {0, -1},
		'E': {0, 1},
	}

	dirCompat = map[rune]rune{
		'N': 'S',
		'S': 'N',
		'W': 'E',
		'E': 'W',
	}

	tRows, tCols int
)

/*
	Find the longest loop starting from "S",
	then return its farthest point(length of the loop divided by 2)

	Algorithm:
	1. Determine the starting point(S)
	2. For every pipe that is connected to S:
	3. Initiate a counter, start BFS, marking the current pipe
	4. Traverse every pipe connected to the current one,
	   adding 1 to the counter
	5. Unmark the current pipe
	6. If S is reached, return the counter number
    7. Otherwise, return 0
	8. Return the biggest counter divided by 2

	Pipe algo: start from one end of a pipe(usually the left/bottom side),
	check whether the direction of the other end matches the opposing side's direction
	[W, E] => [W, E] => [W, S] form a valid path
	[W, E] => [S, E] => [N, S] don't
*/

func main() {
	inputPath := defaultInputPath
	if len(os.Args) == 2 {
		inputPath = os.Args[1]
	}

	input, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("err: could not open %s: %s\n", inputPath, err.Error())
		os.Exit(1)
	}

	parsed := parseInput(input)
	startRow, startCol := findStart(parsed)

	fmt.Printf("Result: %d\n", findMaxCycle(parsed, startRow, startCol)/2)
}

func parseInput(f *os.File) [][]rune {
	scanner := bufio.NewScanner(f)

	res := [][]rune{}

	for scanner.Scan() {
		res = append(res, []rune(scanner.Text()))
	}

	scanner.Text()

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	tCols, tRows = len(res), len(res[0])
	return res
}

func findStart(input [][]rune) (int, int) {
	for i := 0; i < tRows; i++ {
		for j := 0; j < tCols; j++ {
			if input[i][j] == startChar {
				return i, j
			}
		}
	}

	return -1, -1
}

func findMaxCycle(input [][]rune, row, col int) int {
	res := 0

	for _, dir := range directions {
		newRow, newCol := row+dirMap[dir][0], col+dirMap[dir][1]
		if newRow < 0 || newRow >= tRows || newCol < 0 || newCol >= tCols {
			continue
		}

		cycleLen := searchCycle(input, dir, newRow, newCol, 0)
		res = max(res, cycleLen)
	}

	return res
}

func searchCycle(input [][]rune, from rune, row, col, accum int) int {
	cur := input[row][col]

	if cur == '.' || cur == markChar {
		return -1
	}

	if cur == startChar {
		return accum + 1
	}

	res := 0
	if !isConnected(from, pipeMap[cur][0]) && !isConnected(from, pipeMap[cur][1]) {
		return -1
	}

	var newDir rune
	if isConnected(from, pipeMap[cur][0]) {
		newDir = pipeMap[cur][1]
	} else {
		newDir = pipeMap[cur][0]
	}

	newRow, newCol := row+dirMap[newDir][0], col+dirMap[newDir][1]

	if newRow < 0 || newRow >= tRows || newCol < 0 || newCol >= tCols {
		return -1
	}

	input[row][col] = markChar
	res = searchCycle(input, newDir, newRow, newCol, accum+1)
	input[row][col] = cur

	return res
}

func isConnected(from, to rune) bool {
	return dirCompat[from] == to
}
