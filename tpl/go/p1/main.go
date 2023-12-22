package main

import (
	"fmt"
	"os"
)

const (
	defaultInputPath = "../input.txt"
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

	fmt.Printf("input file %s ready to go!\n", input.Name())

	/*
		raw, err := io.ReadAll(input)
		if err != nil {
			fmt.Printf("err: could not read %s file contents: %s", inputPath, err.Error())
			os.Exit(1)
		}
	*/
}
