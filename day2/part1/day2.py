#!/usr/bin/env python3

import collections
import typing

with open("input", "r") as file:
    ids: typing.List[str] = file.read().split("\n")

twice = 0
three = 0

for _id in ids:
    counter = collections.Counter(_id)
    is_twice = False
    is_three = False
    for _, count in counter.items():
        if(count == 2):
            is_twice = True
        elif(count == 3):
            is_three = True
    if(is_twice):
        twice += 1
    if(is_three):
        three += 1

checksum = twice * three

print("checksum: {}".format(checksum))
