package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	numbers, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	var frequency int
	for _, char := range strings.Split(string(numbers), "\n") {
		// fmt.Printf("%v + %v\n", frequency, char)
		frequencyOffset, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		frequency += frequencyOffset
	}
	fmt.Println(frequency)
}
