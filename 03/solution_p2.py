#!/usr/bin/env python3

# Part 2 - lines come in groups of three
# Find a common char for each group and return the total sum

INPUT_FILE = 'input.txt'
GROUP_LEN = 3

def main():
    res = 0
    lines = []
    with open(INPUT_FILE, 'r') as file:
        lines = [line.rstrip() for line in file]

    print(lines)

    res = parse_groups(lines)

    print("Result", res)

def parse_groups(lines) -> int:
    res = 0
    i = 0

    while i < len(lines):
        freqs = [{}, {}, {}]
        j = 0

        for _ in range(GROUP_LEN):
            line = lines[i]
            print(f"Parsing line {line}")

            for ch in line:
                freqs[j][ch] = 1
            i += 1
            j += 1

        for ch in freqs[0].keys():
            if ch in freqs[1] and ch in freqs[2]:
                print(f"Found similar character {ch}, prio {get_priority(ch)}")
                res += get_priority(ch)
                break

    return res

def get_priority(ch) -> int:
    modifier = 0
    if ord(ch) < ord('a'):
        modifier = 38
    else:
        modifier = 96

    return ord(ch) - modifier

main()
