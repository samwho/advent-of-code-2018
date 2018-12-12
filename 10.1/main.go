package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	// position=<-32494,  54541> velocity=< 3, -5>
	inputRegex = regexp.MustCompile(`^position=<\s*([0-9\-]+),\s+([0-9\-]+)>\s+velocity=<\s*([0-9\-]+),\s+([0-9\-]+)>$`)
)

type vec2 struct {
	x, y int
}

func (v vec2) String() string {
	return fmt.Sprintf("vec2{ %d, %d }", v.x, v.y)
}

type point struct {
	position, velocity vec2
}

func (p *point) step() {
	p.position = vec2{
		x: p.position.x + p.velocity.x,
		y: p.position.y + p.velocity.y,
	}
}

func main() {
	fh, _ := os.Open("10.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	points := make([]*point, 0, 300)

	for s.Scan() {
		match := inputRegex.FindStringSubmatch(s.Text())

		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])

		position := vec2{x: px, y: py}
		velocity := vec2{x: vx, y: vy}

		points = append(points, &point{position, velocity})
	}

	for {
		h := heuristic(points)

		if h < 35 {
			print(points)
		}

		step(points)
	}
}

func heuristic(points []*point) int {
	root := points[0]
	sum := 0.0

	for _, point := range points[1:] {
		sum += math.Abs(float64(root.position.x-point.position.x)) +
			math.Abs(float64(root.position.y-point.position.y))
	}

	return int(sum / float64(len(points)))
}

func step(points []*point) {
	for _, point := range points {
		point.step()
	}
}

func print(points []*point) {
	top := points[0].position.y
	left := points[0].position.x
	bottom := points[0].position.y
	right := points[0].position.x

	for _, point := range points {
		if point.position.y < top {
			top = point.position.y
		}
		if point.position.y > bottom {
			bottom = point.position.y
		}
		if point.position.x > right {
			right = point.position.x
		}
		if point.position.x < left {
			left = point.position.x
		}
	}

	width := int(math.Abs(float64(left-right))) + 1
	height := int(math.Abs(float64(top-bottom))) + 1

	grid := make([][]rune, width, width)

	for x := 0; x < width; x++ {
		grid[x] = make([]rune, height, height)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			grid[x][y] = '.'
		}
	}

	for _, point := range points {
		transposed := vec2{
			x: point.position.x - left,
			y: point.position.y - top,
		}

		grid[transposed.x][transposed.y] = '#'
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Print("\n")
	}
}
