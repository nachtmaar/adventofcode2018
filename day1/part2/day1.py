#!/usr/bin/env python3

duplicates = set()
with open('input', 'r') as f:
    frequencies = [int(x) for x in f.read().split("\n")]


def find_frequency():
    sum_frequencies = 0
    index = 0
    while True:
        for frequency in frequencies:
            sum_frequencies += frequency
            previous_length = len(duplicates)
            duplicates.add(sum_frequencies)
            current_length = len(duplicates)
            if(previous_length == current_length):
                print(index)
                print(sum_frequencies)
                return
        index += 1


if __name__ == "__main__":
    find_frequency()
