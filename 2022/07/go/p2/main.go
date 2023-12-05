package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
	For this problem, we should first create a filesystem tree
	from a given input, and then sum all of its nodes' sizes
	in a recursive fashion, caching each dir's size in a hashmap
*/

const (
	inputFile    = "input.txt"
	initialSize  = 70000000
	requiredSize = 30000000
)

// Following the Linux philosophy,
// directories are also files
type file struct {
	name     string
	isDir    bool
	parent   *file
	children map[string]*file
	size     int64
}

func main() {
	// parse the file command by command
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("could not open the input file: %s\n", err.Error())
	}
	defer f.Close()

	tree := parseTree(f)

	// sum all tree elements
	fmt.Println("Result:", sumTree(tree))
}

// Commands to handle - cd *dir*, cd .., ls
func parseTree(f *os.File) *file {
	dummy := newFile("dummy", true, nil)
	curNode := dummy

	fScanner := bufio.NewScanner(f)
	fScanner.Scan()

	for fScanner.Text() != "" {
		tokens := strings.Split(fScanner.Text(), " ")

		switch tokens[1] {
		case "ls":
			for {
				fScanner.Scan()

				line := strings.Split(fScanner.Text(), " ")
				if fScanner.Text() == "" || line[0] == "$" {
					break
				}

				if line[0] == "dir" {
					newNode := newFile(line[1], true, curNode)
					curNode.children[line[1]] = newNode
				} else {
					size, err := strconv.ParseInt(line[0], 10, 64)
					if err != nil {
						fmt.Printf("could not parse line: %s - %s as int64: %s\n", line, line[0], err.Error())
					}

					curNode.size += size
				}
			}
		case "cd":
			if tokens[2] == ".." {
				curNode = curNode.parent
			} else {
				if _, ok := curNode.children[tokens[2]]; !ok {
					newNode := newFile(tokens[2], true, curNode)
					curNode.children[tokens[2]] = newNode
				}

				curNode = curNode.children[tokens[2]]
			}

			fScanner.Scan()
		}
	}

	return dummy
}

func sumTree(root *file) int64 {
	memo := make(map[*file]int64)

	maxHeap := &Int64Heap{}
	heap.Init(maxHeap)

	var sumRec func(*file) int64
	sumRec = func(root *file) int64 {
		if root == nil {
			return 0
		}

		if memo[root] != 0 {
			return memo[root]
		}

		curSize := root.size
		for _, c := range root.children {
			curSize += sumRec(c)
		}

		memo[root] = curSize
		heap.Push(maxHeap, -1*curSize) // -1 to emulate a max heap from a min heap

		return curSize
	}

	usedSize := sumRec(root) // size of the root dir
	reductionTarget := requiredSize - (initialSize - usedSize)

	fmt.Printf("Used size: %d, reduction target: %d\n", usedSize, reductionTarget)

	res := int64(math.MaxInt64)
	fmt.Printf("heap len: %d\n", maxHeap.Len())

	for maxHeap.Len() > 0 {
		cur := -1 * heap.Pop(maxHeap).(int64)

		fmt.Printf("Popped cur: %d\n", cur)

		if cur < reductionTarget {
			break
		}

		res = min(res, cur)
	}

	return res
}

func newFile(name string, isDir bool, parent *file) *file {
	return &file{
		name:     name,
		isDir:    isDir,
		parent:   parent,
		children: make(map[string]*file),
		size:     0,
	}
}
