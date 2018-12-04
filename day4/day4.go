package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"time"
)

type event string

const (
	begins event = "begins"
	awake  event = "awake"
	asleep event = "asleep"
)

type logEntry struct {
	time  time.Time
	guard int
	event event
}

type logEntries []logEntry

func (s logEntries) Len() int {
	return len(s)
}
func (s logEntries) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s logEntries) Less(i, j int) bool {
	return s[i].time.Unix() < s[j].time.Unix()
}

func parseLogEntry(entry string) logEntry {
	var _time string
	var eventString string
	var id int

	re, err := regexp.Compile(`\[(\d+-\d+-\d+ \d+[:]\d+)\] (.+)`)
	if err != nil {
		panic(err)
	}
	var matches = re.FindStringSubmatch(entry)
	_time = matches[1]
	eventString = matches[2]

	t, err := time.Parse("2006-01-02 15:04", _time)
	if err != nil {
		panic(err)
	}
	// check if entry contains id
	fmt.Sscanf(eventString, "Guard #%d", &id)
	// determine one of the event types
	var _event event
	if strings.Contains(eventString, "begins shift") {
		_event = begins
	} else if strings.Contains(eventString, "wakes up") {
		_event = awake
	} else if strings.Contains(eventString, "falls asleep") {
		_event = asleep
	} else {
		panic("unknown event type")
	}

	return logEntry{t, id, _event}
}

func findGuardMostFrequentlyAsleepSameMinute(_logEntries logEntries) (int, int) {
	// guardMinuteFrequency[guard][minute] = frequency

	var guardMinuteFrequency = make(map[int]map[int]int)
	var currentGuard int
	var timeAsleep time.Time
	for _, _logEntry := range _logEntries {

		switch _logEntry.event {
		case asleep:
			timeAsleep = _logEntry.time
			break
		case awake:
			for minute := int(timeAsleep.Minute()); minute < int(_logEntry.time.Minute()); minute++ {
				if guardMinuteFrequency[currentGuard] == nil {
					guardMinuteFrequency[currentGuard] = make(map[int]int)
				}
				guardMinuteFrequency[currentGuard][minute]++
			}
			break
		case begins:
			currentGuard = _logEntry.guard
			break
		}
	}
	var guard int
	var minute int
	var max int
	for _guard, minuteFrequency := range guardMinuteFrequency {
		for _minet, frequency := range minuteFrequency {
			if frequency > max {
				max = frequency
				guard = _guard
				minute = _minet
			}
		}
	}
	return guard, minute
}

func findSleepiestMinute(_logEntries logEntries, guard int) int {
	// remember the frequency for each minute the guard slept
	var sleeps = make(map[int]int, 10)
	var currentGuard int
	var timeAsleep time.Time
	for _, _logEntry := range _logEntries {

		switch _logEntry.event {
		case asleep:
			if currentGuard != guard {
				continue
			}
			timeAsleep = _logEntry.time
			break
		case awake:
			if currentGuard != guard {
				continue
			}
			for minute := int(timeAsleep.Minute()); minute < int(_logEntry.time.Minute()); minute++ {
				sleeps[minute]++
			}
			break
		case begins:
			currentGuard = _logEntry.guard
			break
		}
	}

	var sleepiestMinute int
	var max int
	for minute, count := range sleeps {
		if count > max {
			max = count
			sleepiestMinute = minute
		}
	}

	return sleepiestMinute
}

func findSleepiestGuard(_logEntries logEntries) int {
	// store for each guard how many minutes he slept
	var res = make(map[int]int)
	var currentGuard int
	var timeAsleep time.Time
	for _, _logEntry := range _logEntries {
		switch _logEntry.event {
		case asleep:
			timeAsleep = _logEntry.time
			break
		case awake:
			var slept = int(_logEntry.time.Sub(timeAsleep).Minutes())
			if slept <= 0 {
				panic("negative sleep time")
			}
			fmt.Printf("%v => %v = %v\n", timeAsleep.UTC(), _logEntry.time.UTC(), slept)
			res[currentGuard] += slept
			break
		case begins:
			currentGuard = _logEntry.guard
			break
		}
	}

	var max int
	var sleepiestGuard int
	for guard, minutes := range res {
		if minutes > max {
			max = minutes
			sleepiestGuard = guard
		}
	}
	fmt.Printf("max minutes slept: %v\n", max)
	return sleepiestGuard
}

func main() {
	bytes, error := ioutil.ReadFile("input")
	if error != nil {
		panic(error)
	}
	var input = strings.Split(string(bytes), "\n")
	var _logEntries []logEntry
	fmt.Println(len(input))
	for _, line := range input {
		_logEntries = append(_logEntries, parseLogEntry(line))
	}
	sort.Sort(logEntries(_logEntries))
	for _, logEntry := range _logEntries {
		fmt.Printf("%v %v, %v\n", logEntry.time, logEntry.guard, logEntry.event)
	}
	var sleepiestGuard = findSleepiestGuard(_logEntries)
	var sleepiestMinute = findSleepiestMinute(_logEntries, sleepiestGuard)
	guard, minute := findGuardMostFrequentlyAsleepSameMinute(_logEntries)

	fmt.Printf("max asleep: %v at %v\n", sleepiestGuard, sleepiestMinute)
	fmt.Printf("part1: %v\n", sleepiestGuard*sleepiestMinute)
	fmt.Println(guard, minute)
	fmt.Printf("part2: %v\n", guard*minute)
}
