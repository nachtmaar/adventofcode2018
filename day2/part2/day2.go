package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func wordEquality(id string, id2 string) int {
	var equality int
	for idx := range id {
		if id[idx] == id2[idx] {
			equality++
		}
	}
	return equality
}

func getEquality(id string, id2 string) string {
	var equality = ""
	for idx, char := range id {
		if id[idx] == id2[idx] {
			equality += string(char)
		}
	}
	return equality
}

func main() {
	text, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	var ids = strings.Split(string(text), "\n")

	for idx, id := range ids {
		for _, id2 := range ids[idx:] {
			var wordLength = len(id)
			if wordLength != len(id2) {
				panic("words are expected to have same length")
			}
			if wordEquality(id, id2) == wordLength-1 {
				fmt.Println(id)
				fmt.Println(id2)
				fmt.Println(getEquality(id, id2))
			}
		}
	}
}
