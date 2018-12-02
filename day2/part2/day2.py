#!/usr/bin/env python3

import collections
import typing


def equality(word: str, word2: str) -> int:
    assert len(word) == len(word2), 'words do not have equal length'
    equality = 0
    for i in range(len(word)):
        if(word[i] == word2[i]):
            equality += 1
    return equality


def find_box_ids(ids) -> typing.Tuple[str, str]:
    word_length = len(ids[0])
    for idx, _id in enumerate(ids):
        for _id2 in ids[idx:]:
            if(_id == _id2):
                continue
            word_equality = equality(_id, _id2)
            if(word_equality == word_length - 1):
                return _id, _id2


if __name__ == "__main__":
    with open("input", "r") as file:
        ids: typing.List[str] = file.read().split("\n")
    _id1, _id2 = find_box_ids(ids)
    print(_id1, _id2)
    res = ""
    for i in range(len(_id1)):
        if(_id1[i] == _id2[i]):
            res += _id1[i]
    print(res)
