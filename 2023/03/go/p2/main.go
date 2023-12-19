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
	1. Detect numbers the same way as in p1,
	   except now the only valid part for detection is '*'
	2. For every detected number, save the asterisk location in a map as key
	   and add a part number array as value
	3. Filter out asterisks with two neighboring numbers only
	4. Get their parts' products sums
*/

const (
	defaultInputPath = "input.txt"
)

type Gear struct {
	row int
	col int
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

	raw, err := io.ReadAll(input)
	if err != nil {
		fmt.Printf("err: could not read %s file contents: %s", inputPath, err.Error())
		os.Exit(1)
	}

	if raw[len(raw)-1] == '\n' {
		raw = raw[:len(raw)-1]
	}

	fmt.Printf("Answer: %d\n", extractGears(string(raw)))
}

func extractGears(raw string) int {
	matrix := strings.Split(raw, "\n")

	gearNumbers := make(map[Gear][]int)
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		numStart, numEnd := -1, -1

		for j := 0; j < cols; j++ {
			if isDigit(matrix[i][j]) {
				numStart = j

				for j < cols && isDigit(matrix[i][j]) {
					j++
				}

				numEnd = j - 1

				gears := findGears(matrix, i, numStart, numEnd)
				if len(gears) > 0 {
					num, _ := strconv.Atoi(matrix[i][numStart : numEnd+1])

					for _, gear := range gears {
						gearNumbers[gear] = append(gearNumbers[gear], num)
					}
				}

				numStart, numEnd = -1, -1
			}
		}
	}

	res := 0

	for _, nums := range gearNumbers {
		if len(nums) == 2 {
			res += nums[0] * nums[1]
		}
	}

	return res
}

func findGears(matrix []string, row, start, end int) []Gear {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
	gearSet := map[Gear]bool{}

	rows, cols := len(matrix), len(matrix[0])

	for col := start; col <= end; col++ {
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				if matrix[newRow][newCol] == '*' {
					gearSet[Gear{row: newRow, col: newCol}] = true
				}
			}
		}
	}

	gears := make([]Gear, 0, len(gearSet))
	for g := range gearSet {
		gears = append(gears, g)
	}

	return gears
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
