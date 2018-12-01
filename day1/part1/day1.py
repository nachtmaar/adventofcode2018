#!/usr/bin/env python3
import functools

with open('input', 'r') as f:
    lines = [int(x) for x in f.read().split("\n")]
    frequency = functools.reduce(lambda x,y: x+y, lines)
    print(frequency)

