#!/usr/bin/env python3

INPUT_FILE="input.txt"

beat_map = {
    "A": "C",
    "B": "A",
    "C": "B"
}

beaten_map = {
    "A": "B",
    "B": "C",
    "C": "A"
}

score_map = {
    "A": 1,
    "B": 2,
    "C": 3
}


def main():
    score = 0
    with open(INPUT_FILE) as file:
        for line in file:
            symbol, target = line.split()

            score += calc_outcome(symbol, target)

    print("Score:", score)


def calc_outcome(symbol, target) -> int:
    match target:
        case "X": # lose
            return score_map[beat_map[symbol]]

        case "Y": # draw
            return 3 + score_map[symbol]

        case "Z": # win
            return 6 + score_map[beaten_map[symbol]]

        case _:
            raise Exception("invalid target")

main()
