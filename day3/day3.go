package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type claim struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

type point struct {
	x int
	y int
}

func main() {
	bytes, error := ioutil.ReadFile("input")
	if error != nil {
		panic(error)
	}
	var textClaims = strings.Split(string(bytes), "\n")
	var claims []claim
	for _, textClaim := range textClaims {
		var _claim = readClaim(textClaim)
		claims = append(claims, _claim)
	}
	overlapses, intactClaim := countOverlapes(claims)

	fmt.Printf("number of overlapses: %v\n", overlapses)
	fmt.Printf("intact: %v\n", intactClaim)
}

func countOverlapes(claims []claim) (int, claim) {
	// store for each point how often they have been claimed
	var frequency = make(map[point]int)
	var claimMap = make(map[point]claim)
	// run over claims
	for _, claim := range claims {
		// run over all points in claim
		for x := claim.left; x < claim.left+claim.width; x++ {
			for y := claim.top; y < claim.top+claim.height; y++ {
				var point = point{x, y}
				frequency[point]++
				claimMap[point] = claim
			}
		}
	}
	// determine number of overlapses
	var count int
	for _, _count := range frequency {
		if _count > 1 {
			count++
		}
	}

	// determine claim that has no overlapse at all
	var intactClaim claim
	for _, claim := range claims {
		// run over all points in claim
		for x := claim.left; x < claim.left+claim.width; x++ {
			for y := claim.top; y < claim.top+claim.height; y++ {
				var point = point{x, y}
				// intact as long as no overlapse found
				if frequency[point] != 1 {
					goto start
				}
			}
		}
		intactClaim = claim
		break
	start:
	}

	return count, intactClaim
}

func readClaim(_claim string) claim {
	var pattern = `#(?P<id>\d+) @ (?P<left>\d+),(?P<top>\d+)[:] (?P<width>\d+)x(?P<height>\d+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(_claim)

	captures := make(map[string]string)
	for index, name := range re.SubexpNames() {
		if index != 0 && name != "" {
			captures[name] = matches[index]
		}
	}
	left, error := strconv.Atoi(captures["left"])
	if error != nil {
		panic(error)
	}
	top, error := strconv.Atoi(captures["top"])
	if error != nil {
		panic(error)
	}
	width, error := strconv.Atoi(captures["width"])
	if error != nil {
		panic(error)
	}
	height, error := strconv.Atoi(captures["height"])
	if error != nil {
		panic(error)
	}

	return claim{
		id:     captures["id"],
		left:   left,
		top:    top,
		width:  width,
		height: height,
	}
}
