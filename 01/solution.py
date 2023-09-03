#!/usr/bin/env python3

# Task
# There is a number of elves each carrying
# a set amount of calories, with each elves'
# load being separated by a newline
# We need to find out the biggest amount of calories
# carried by a given elf

# Solution
# Initialize a max-heap
# Read the file line by line keeping a rolling sum between newlines
# At the newline, push the sum to a max heap
# Upon finishing the file read, return the top of max heap 

from sys import argv
from heapq import heappush, heappop

TOP_ELVES_COUNT = 3

def main():
    if len(argv) != 2:
        raise error('Usage: solution.py *filename*')

    cur_sum = 0
    heap = list()
    filename = argv[1]
    with open(filename) as file:
        for line in file:
            if line == '\n':
                push_elf_calories(heap, cur_sum)
                cur_sum = 0
            else:
                cur_sum += int(line)

    res = 0
    print(f"heap len: {len(heap)}")
    while heap:
        res += heappop(heap) 

    print(f"Top {TOP_ELVES_COUNT} elves carry {res} calories")

def push_elf_calories(heap, calories):
    if len(heap) < TOP_ELVES_COUNT:
        heappush(heap, calories)
    elif heap[0] < calories:
        heappop(heap)
        heappush(heap, calories)

main()
