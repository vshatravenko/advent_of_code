from collections import deque
import re

INPUT_FILE = 'input.txt'

# read all the input file lines
def main():
    lines = []
    with open(INPUT_FILE) as f:
        lines = f.readlines()

    actions_start = find_actions_start(lines)

    stacks = parse_stacks(lines[:actions_start])

    process_actions(lines[actions_start+1:], stacks)
    print(stacks)
  
    print(extract_res(stacks))

def find_actions_start(lines) -> int:
    i = 0
    while lines[i] != '\n':
        i += 1

    return i

# Parse the stacks
def parse_stacks(lines) -> list[deque]:
    # -3 because of ' \n'
    count = int(lines[len(lines)-1][-3])

    stacks = [deque() for _ in range(count)]
    # Each stack elem takes 4 cells except for the last one
    col_step = 4

    for i in range(len(lines)-2, -1, -1):
        j = 1 
        while j < len(lines[i]):
            if lines[i][j] != ' ': 
                idx = (j - 1) // 4
                stacks[idx].append(lines[i][j])
            j += col_step

    return stacks

# Perform actions line by line
def process_actions(lines, stacks):
    pattern = re.compile('move (\d*) from (\d*) to (\d*)')
    for line in lines:
        m = pattern.match(line)

        do_action(stacks, int(m.group(1)), int(m.group(2))-1, int(m.group(3))-1)

def do_action(stacks, count, from_idx, to_idx):
    print(f"pushing {count} times from {from_idx} to {to_idx}")
    items = deque()
    for _ in range(count):
        item = stacks[from_idx].pop()
        items.append(item)

    for _ in range(count):
        item = items.pop()
        stacks[to_idx].append(item)

def extract_res(stacks) -> str:
    res = [stack[-1] for stack in stacks]

    return "".join(res)


if __name__ == "__main__":
    main()
