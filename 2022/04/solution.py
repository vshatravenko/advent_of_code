#!/usr/bin/env python3

INPUT_FILE="input.txt"

def main():
    res = 0
    with open(INPUT_FILE) as file:
        for line in file:
            pair1, pair2 = line.split(',')
            print(f"Checking {pair1} & {pair2}")
            if check_overlap(pair1, pair2):
                res += 1
            else:
                print("Rejected", pair1, pair2)
    print("Res:", res)

def check_overlap(p1, p2) -> bool:
    p1_start, p1_end = p1.split('-')
    p2_start, p2_end = p2.split('-')

    if int(p1_start) >= int(p2_start) and int(p1_start) <= int(p2_end):
        return True
    elif int(p1_end) >= int(p2_start) and int(p1_end) <= int(p2_end):
        return True
    elif int(p2_start) >= int(p1_start) and int(p2_start) <= int(p1_end):
        return True
    elif int(p2_end) >= int(p1_start) and int(p2_end) <= int(p1_end):
        return True

    return False

main()
