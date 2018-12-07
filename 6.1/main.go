package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

func main() {
	fh, _ := os.Open("6.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	var grid [400][400]int
	var coordsSize [400]int

	coords := make([]coordinate, 0, 50)

	for s.Scan() {
		parts := strings.Split(s.Text(), ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coords = append(coords, coordinate{x, y})
	}

	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			grid[x][y] = -1
		}
	}

	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			minDistance := 10000000.0
			minDistanceIdx := -1
			unique := false

			for coordIdx, coord := range coords {
				distance := math.Abs(float64(x-coord.x)) + math.Abs(float64(y-coord.y))
				if distance < minDistance {
					minDistance = distance
					minDistanceIdx = coordIdx
					unique = true
				} else if distance == minDistance {
					unique = false
				}
			}

			if unique {
				grid[x][y] = minDistanceIdx
				coordsSize[minDistanceIdx]++
			}
		}
	}

	for x := 0; x < 400; x++ {
		coord := grid[x][0]
		if coord > -1 {
			coordsSize[coord] = -1
		}
		coord = grid[x][399]
		if coord > -1 {
			coordsSize[coord] = -1
		}
	}

	for y := 0; y < 400; y++ {
		coord := grid[0][y]
		if coord > -1 {
			coordsSize[coord] = -1
		}
		coord = grid[399][y]
		if coord > -1 {
			coordsSize[coord] = -1
		}
	}

	maxSize := -1
	maxIdx := -1
	for idx, size := range coordsSize {
		if size > maxSize {
			maxSize = size
			maxIdx = idx
		}
	}

	fmt.Printf("%d, size: %d\n", maxIdx, maxSize)
}
