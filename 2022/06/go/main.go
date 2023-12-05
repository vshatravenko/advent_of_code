package main

import (
	"fmt"
	"os"
)

const (
	inputFile  = "input.txt"
	uniqWindow = 14
)

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the input file: %s\n", err.Error())
		os.Exit(1)
	}

	pos := findMarker(data)

	if pos == -1 {
		fmt.Println("Could not find the marker!")
	} else {
		fmt.Printf("Result: pos - %d; substr - %s\n", pos, string(data[pos-uniqWindow:pos])) // offset by 1 due to the 0-based index
	}
}

func findMarker(data []byte) int {
	set := make(map[byte]int)
	uniqCount := 0

	start := 0

	for i := start; i < uniqWindow; i++ {
		_, ok := set[data[i]]
		if !ok {
			uniqCount++
		}

		set[data[i]]++
	}

	for i := uniqWindow; i < len(data); i++ {
		if uniqCount == uniqWindow {
			return i
		}

		if set[data[start]] == 1 {
			uniqCount--
		}
		set[data[start]]--

		if val, ok := set[data[i]]; !ok || val == 0 {
			uniqCount++
		}
		set[data[i]]++

		start++
	}

	return -1
}
