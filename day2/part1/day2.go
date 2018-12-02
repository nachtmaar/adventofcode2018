package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	text, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	var ids = strings.Split(string(text), "\n")

	var twice int
	var three int
	for _, id := range ids {
		var letterFrequencies = countLetters(id)
		var isTwice = false
		var isThree = false
		for _, frequency := range letterFrequencies {
			if frequency == 2 {
				isTwice = true
			} else if frequency == 3 {
				isThree = true
			}
		}
		if isTwice {
			twice++
		}
		if isThree {
			three++
		}
	}
	var checksum = twice * three
	fmt.Printf("checksum: %v\n", checksum)
}

func countLetters(word string) map[rune]int {
	var res = make(map[rune]int)
	for _, char := range word {
		res[char]++
	}
	for k, v := range res {
		fmt.Printf("%v = %v\n", k, v)
	}
	return res
}
