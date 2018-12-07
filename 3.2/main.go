package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	regex = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
)

func main() {
	fh, _ := os.Open("3.2/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)
	ss := make([]string, 0, 1500)

	var fabric [1000][1000]int

	for s.Scan() {
		ss = append(ss, s.Text())
	}

	for _, input := range ss {
		_, left, top, width, height := parseInput(input)

		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				fabric[x][y]++
			}
		}
	}

	for _, input := range ss {
		id, left, top, width, height := parseInput(input)

		isgood := true
		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				if fabric[x][y] != 1 {
					isgood = false
					break
				}
			}

			if !isgood {
				break
			}
		}

		if isgood {
			fmt.Printf("%d\n", id)
		}
	}
}

func parseInput(s string) (int, int, int, int, int) {
	match := regex.FindStringSubmatch(s)

	id, _ := strconv.Atoi(match[1])
	left, _ := strconv.Atoi(match[2])
	top, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])

	return id, left, top, width, height
}
