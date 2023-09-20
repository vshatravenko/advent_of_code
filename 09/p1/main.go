package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

	The goal is to simulate the movement of a rope's tail and count all the unique positions it visits

	The main parts of the task are:
	1. Tail movement simulation(especially diagonal movement)
	2. Unique positions filtering

	Filtering is achieved by keeping a Set of (X, Y) positions

	Algo:
	1. Prepare mappings for L, R, D, U directions
	2. Movement simulation:
       First evaluate the head movement, then the tail one
	   If there is an overlap between the head and the tail or they're adjacent, skip
       If tail is in the same column/row, move straight
	   Else, move diagonally(apply two motions, e.g. L and U)
	   Record the new position
*/

const (
	defaultInputFile = "input.txt"
)

var (
	motions = map[string][]int{
		"L": {0, -1},
		"R": {0, 1},
		"U": {1, 0},
		"D": {-1, 0},
	}
)

type position struct {
	x, y int
}

func main() {
	inputFile := defaultInputFile
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("could not open %s: %s\n", inputFile, err)
		os.Exit(1)
	}

	visited := parseSim(f)

	fmt.Printf("Log: %v\nResult: %d\n", visited, len(visited))

}

func parseSim(f *os.File) map[position]bool {
	head, tail := []int{0, 0}, []int{0, 0} // starting position
	visited := map[position]bool{{0, 0}: true}

	fScanner := bufio.NewScanner(f)

	for fScanner.Scan() {
		tokens := strings.Split(fScanner.Text(), " ")

		direction := tokens[0]
		distance, err := strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Printf("Could not parse line %s: %s\n", tokens, err)
			os.Exit(1)
		}

		fmt.Printf("Moving head to %s %d times\n", direction, distance)
		for i := 0; i < distance; i++ {
			head[0] += motions[direction][0]
			head[1] += motions[direction][1]

			if !isAdjacent(head, tail) {
				moveTail(head, tail)
				fmt.Printf("Moved tail to %v\n", tail)

				curPos := position{tail[0], tail[1]}

				if _, ok := visited[curPos]; !ok {
					visited[curPos] = true
				}
			}
		}
	}

	return visited
}

func moveTail(head, tail []int) {
	for i := 0; i < len(head); i++ {
		if head[i] > tail[i] {
			tail[i]++
		} else if head[i] < tail[i] {
			tail[i]--
		}
	}
}

// Also returns true when there is an overlap
func isAdjacent(head, tail []int) bool {
	colDiff := abs(max(head[0], tail[0]) - min(head[0], tail[0]))
	rowDiff := abs(max(head[1], tail[1]) - min(head[1], tail[1]))

	fmt.Printf("head pos: %v; tail pos: %v, colDiff: %d, rowDiff: %d, adjacent: %v\n", head, tail, colDiff, rowDiff, (colDiff <= 1 && rowDiff <= 1))

	return colDiff <= 1 && rowDiff <= 1
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
