package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// store all summed frequencies in a map with their occurences
	var storedFrequencies = make(map[int]int)
	// contains the sum of all frequencies
	var sumFrequency int

	// read frequencies from file
	numbers, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// convert frequencies to string array
	var inputFrequencies = strings.Split(string(numbers), "\n")
	// convert frequencies to int array
	var inputFrequenciesInt []int
	for _, element := range inputFrequencies {
		frequency, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		inputFrequenciesInt = append(inputFrequenciesInt, frequency)
	}

	// several runs over frequencies is required to find duplicates
	for {
		for _, frequency := range inputFrequenciesInt {
			sumFrequency += frequency
			_, ok := storedFrequencies[sumFrequency]
			if !ok {
				storedFrequencies[sumFrequency] = 0
			}
			storedFrequencies[sumFrequency] = storedFrequencies[sumFrequency] + 1
			if storedFrequencies[sumFrequency] > 1 {
				fmt.Printf("duplicate frequency found: %v\n", sumFrequency)
				return
			}
		}
	}

}
