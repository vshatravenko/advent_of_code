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
	However, now the rope consists of 10 knots, with the knot movement logic preserved

	The main parts of the task are:
	1. Tail movement simulation(especially diagonal movement)
	2. Unique positions filtering

	Filtering is achieved by keeping a Set of (X, Y) positions

	Algo:
	1. Prepare mappings for L, R, D, U directions
	2. Movement simulation:
	   Evaluate movement for each knot, comparing it to the previous one
	   If there is an overlap between the current and the previous knot or they're adjacent, skip
       If the current knot is in the same column/row, move straight
	   Else, move diagonally(apply two motions, e.g. L and U)
	   If this is the tail(10th knot), record the new position
*/

const (
	defaultInputFile = "input.txt"
	ropeLen          = 10
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

	visited := parseSim(f, ropeLen)

	fmt.Printf("Result: %d\n", len(visited))

}

func parseSim(f *os.File, ropeLen int) map[position]bool {
	knots := make([][]int, ropeLen)
	for i := 0; i < ropeLen; i++ {
		knots[i] = []int{0, 0}
	}
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
			knots[0][0] += motions[direction][0]
			knots[0][1] += motions[direction][1]

			for i := 1; i < ropeLen; i++ {
				if !isAdjacent(knots[i-1], knots[i]) {
					moveKnot(knots[i-1], knots[i])
					fmt.Printf("Moved knot %d to %v\n", i, knots[i])

					if i == ropeLen-1 {
						curPos := position{knots[i][0], knots[i][1]}

						if _, ok := visited[curPos]; !ok {
							visited[curPos] = true
						}
					}
				}

			}
		}
	}

	return visited
}

func moveKnot(prev, cur []int) {
	for i := 0; i < 2; i++ { // x and y
		if prev[i] > cur[i] {
			cur[i]++
		} else if prev[i] < cur[i] {
			cur[i]--
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
