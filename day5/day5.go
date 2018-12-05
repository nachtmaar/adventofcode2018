package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type polymers []unit
type polarity string

const (
	positive = "positive"
	negative = "negative"
)

type unit struct {
	name     string
	polarity polarity
}

func parseUnit(input rune) unit {
	var _polarity polarity
	if unicode.IsLower(input) {
		_polarity = negative
	} else {
		_polarity = positive
	}
	return unit{string(input), _polarity}
}

func fmtPolymer(_polymers polymers) string {
	var res = ""
	for _, _unit := range _polymers {
		res += _unit.name
	}
	return res
}

func reacts(a unit, b unit) bool {
	return strings.ToLower(a.name) == strings.ToLower(b.name) && a.polarity != b.polarity
}

func react(_polymers polymers) polymers {
	var res polymers
	var length = len(_polymers)
	for index := 0; index < length; index++ {

		var _unitA = _polymers[index]
		if index < length-1 {
			var _unitB = _polymers[index+1]
			if !reacts(_unitA, _unitB) {
				res = append(res, _unitA)
			} else {
				// we have to drop _unitB
				index++
			}
		} else {
			res = append(res, _unitA)
		}

	}
	return res
}

func reducedSize(_polymers polymers) int {
	var doesReact = true
	for doesReact {
		var lengthBefore = len(_polymers)
		_polymers = react(_polymers)
		var lengthAfter = len(_polymers)
		doesReact = lengthBefore != lengthAfter
		// fmt.Printf("%v => %v\n", lengthBefore, lengthAfter)
	}
	return len(_polymers)
}

func minReducedSize(_polymers polymers) int {
	var uniqueUnits = make(map[string]unit)
	for _, _unit := range _polymers {
		// build set of unique units
		uniqueUnits[strings.ToLower(_unit.name)] = _unit
	}
	fmt.Println(uniqueUnits)
	var min = -1
	for _, _unit := range uniqueUnits {
		var polymerswithoutUnit = without(_polymers, _unit)
		var _reducedSize = reducedSize(polymerswithoutUnit)
		fmt.Printf("%v => %v\n", strings.ToLower(_unit.name), _reducedSize)
		if min == -1 {
			min = _reducedSize
		} else if _reducedSize < min {
			min = _reducedSize
		}

	}
	return min
}

func without(_polymers polymers, _unit unit) polymers {
	var res polymers
	for _, _unit2 := range _polymers {
		if strings.ToLower(_unit.name) == strings.ToLower(_unit2.name) {
			continue
		}
		res = append(res, _unit2)
	}
	return res
}

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	var _polymers polymers

	// read input into polymers type
	for _, unitByte := range bytes {
		_polymers = append(_polymers, parseUnit(rune(unitByte)))
	}

	fmt.Printf("part1: %v\n", reducedSize(_polymers))
	fmt.Printf("part2: %v\n", minReducedSize(_polymers))
}
