package main

import (
	"bufio"
	"fmt"
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
	inputFile = "input.txt"
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
	var res int64
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
		if curSize <= 100000 {
			res += curSize
		}

		return curSize
	}

	sumRec(root)

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
