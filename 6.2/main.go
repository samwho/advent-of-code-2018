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
	fh, _ := os.Open("6.2/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	coords := make([]coordinate, 0, 50)

	for s.Scan() {
		parts := strings.Split(s.Text(), ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coords = append(coords, coordinate{x, y})
	}

	area := 0

	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			sumDistance := 0.0
			for _, coord := range coords {
				distance := math.Abs(float64(x-coord.x)) + math.Abs(float64(y-coord.y))
				sumDistance += distance
			}

			if sumDistance < 10000.0 {
				area++
			}
		}
	}

	fmt.Printf("%d\n", area)
}
