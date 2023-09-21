package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultInputPath = "input.txt"
	screenWidth      = 40
	screenHeight     = 6
)

type Op struct {
	name   string
	value  int
	cycles int
}

func main() {
	inputPath := defaultInputPath
	if len(os.Args) == 2 {
		inputPath = os.Args[1]
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Could not open %s: %s\n", inputPath, err)
		os.Exit(1)
	}

	// parse input commands
	ops := parseInput(inputFile)
	screen := initScreen(screenWidth, screenHeight)

	// print pixels on the CRT screen
	res := drawPixels(screen, ops)

	printPixels(res)
}

func parseInput(f *os.File) []*Op {
	res := []*Op{}
	fScanner := bufio.NewScanner(f)

	for fScanner.Scan() {
		tokens := strings.Split(fScanner.Text(), " ")

		if len(tokens) == 1 {
			res = append(res, &Op{name: "noop", cycles: 1})
		} else {
			value, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Printf("err: could not parse line %s: %s\n", fScanner.Text(), err)
				os.Exit(1)
			}

			res = append(res, &Op{name: tokens[0], value: value, cycles: 2})
		}
	}

	return res
}

/*
Returns a drawn image on the CRT screen
The screen dimensions are 40 * 6
The image is drawn one pixel per cycle
e.g. 41st cycle draws the first pixel of the second row

The register value contains the location of the middle of a sprite which is 3 pixels wide
If the sprite is overlapping the currently drawn pixel,
we should draw it as "#", otherwise we put a "."
*/
func drawPixels(screen [][]string, ops []*Op) [][]string {
	opCount := len(ops)
	cycleCount := len(screen) * len(screen[0])
	registerVal := 1
	opIdx := 0
	pendingVal := 0
	isPendingAddx := false

	for cycle := 1; cycle <= cycleCount && opIdx < opCount; cycle++ {
		ops[opIdx].cycles--

		if isPendingAddx {
			registerVal += pendingVal
			isPendingAddx = false
		}

		if ops[opIdx].cycles == 0 {
			if ops[opIdx].name == "addx" {
				pendingVal = ops[opIdx].value
				isPendingAddx = true
			}

			opIdx++
		}

		drawPixel(screen, len(screen), len(screen[0]), cycle, registerVal)
	}

	return screen
}

func drawPixel(screen [][]string, height, width, cycle, spriteMid int) {
	var pixelX, pixelY int
	if cycle > width {
		if cycle%width == 0 {
			pixelX = width - 1
		} else {
			pixelX = cycle%width - 1
		}

		pixelY = (cycle - 1) / width
		if pixelY == height {
			pixelY--
		}
	} else {
		pixelX, pixelY = cycle-1, 0
	}

	for _, offset := range []int{-1, 0, 1} {
		if spriteMid+offset == pixelX {
			screen[pixelY][pixelX] = "#"
			return
		}
	}

	screen[pixelY][pixelX] = "."
}

func initScreen(width, height int) [][]string {
	screen := make([][]string, screenHeight)
	for i := 0; i < screenHeight; i++ {
		screen[i] = make([]string, screenWidth)
	}

	return screen
}

func printPixels(pixels [][]string) {
	for _, row := range pixels {
		fmt.Println(strings.Join(row, ""))
	}
}
