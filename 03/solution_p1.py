#!/usr/bin/env python3


# Task - find all shared characters for every line(rucksack)
# and return their ASCII value sum
#
# Solution
# For every line, separate the string in two halves,
# initialize two dicts, and add each half's chars there
# Then, find matching chars by traversing each dict

INPUT_FILE = 'input.txt'

def main():
    res = 0
    with open(INPUT_FILE) as file:
        for line in file:
            res += parse_line(line)

    print("Result", res)

def parse_line(line) -> int:
    print("Parsing line", line)
    cutoff = len(line) // 2
    lo, hi = line[:cutoff], line[cutoff:]
    lo_freq, hi_freq = {}, {}

    i = 0
    while i < cutoff:
        lo_freq[lo[i]] = lo_freq.get(i, 0) + 1
        hi_freq[hi[i]] = hi_freq.get(i, 0) + 1
        i += 1

    res = 0
    for ch in lo_freq.keys():
        if ch in hi_freq:
            freq = min(lo_freq[ch], hi_freq[ch])
            res += get_priority(ch) * freq
            print(f"Found {freq} similar {ch} chars, adding {get_priority(ch)}")

    return res

def get_priority(ch) -> int:
    modifier = 0
    if ord(ch) < ord('a'):
        modifier = 38
    else:
        modifier = 96

    return ord(ch) - modifier

main()
