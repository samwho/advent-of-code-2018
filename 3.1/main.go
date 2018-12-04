package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	regex = regexp.MustCompile(`^#\d+ @ (\d+),(\d+): (\d+)x(\d+)$`)
)

func main() {
	fh, _ := os.Open("3.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	var fabric [1000][1000]int

	for s.Scan() {
		left, top, width, height := parseInput(s.Text())

		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				fabric[x][y]++
			}
		}
	}

	sum := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if fabric[x][y] > 1 {
				sum++
			}
		}
	}

	fmt.Printf("%d\n", sum)
}

func parseInput(s string) (int, int, int, int) {
	match := regex.FindStringSubmatch(s)

	left, _ := strconv.Atoi(match[1])
	top, _ := strconv.Atoi(match[2])
	width, _ := strconv.Atoi(match[3])
	height, _ := strconv.Atoi(match[4])

	return left, top, width, height
}
