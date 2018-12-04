package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var (
	shiftRegex = regexp.MustCompile(`^\[(.+)\] Guard #(\d+) begins shift$`)
	wakeRegex  = regexp.MustCompile(`^\[(.+)\] wakes up$`)
	sleepRegex = regexp.MustCompile(`^\[(.+)\] falls asleep$`)
)

func main() {
	fh, _ := os.Open("4.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)
	ss := make([]string, 0, 1500)

	for s.Scan() {
		ss = append(ss, s.Text())
	}

	sort.Strings(ss)

	m := make(map[string]map[int]int)
	var guard string
	var minuteAsleep int

	for _, event := range ss {
		if shiftRegex.MatchString(event) {
			match := shiftRegex.FindStringSubmatch(event)
			guard = match[2]
		} else if sleepRegex.MatchString(event) {
			match := sleepRegex.FindStringSubmatch(event)
			minuteAsleep, _ = strconv.Atoi(match[1][len(match[1])-2:])
		} else if wakeRegex.MatchString(event) {
			match := wakeRegex.FindStringSubmatch(event)
			minuteAwake, _ := strconv.Atoi(match[1][len(match[1])-2:])

			for i := minuteAsleep; i < minuteAwake; i++ {
				_, ok := m[guard]
				if !ok {
					m[guard] = make(map[int]int)
				}

				m[guard][i]++
			}
		} else {
			panic(event)
		}
	}

	maxGuardId := ""
	maxGuardMinutes := -1

	for guardId, minutes := range m {
		sum := 0
		for _, v := range minutes {
			sum += v
		}

		if sum > maxGuardMinutes {
			maxGuardId = guardId
			maxGuardMinutes = sum
		}
	}

	maxGuardIdInt, _ := strconv.Atoi(maxGuardId)

	maxMinute := -1
	maxMinuteCount := -1
	for minute, timesAsleepThisMinute := range m[maxGuardId] {
		if timesAsleepThisMinute > maxMinuteCount {
			maxMinute = minute
			maxMinuteCount = timesAsleepThisMinute
		}
	}

	fmt.Printf("%d\n", maxGuardIdInt*maxMinute)
}
