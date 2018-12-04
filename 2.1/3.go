package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fh, _ := os.Open("3/3.data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	sumTwos := 0
	sumThrees := 0

	for s.Scan() {
		twos, threes := countDupsTrips(s.Text())
		sumTwos += twos
		sumThrees += threes
	}

	fmt.Printf("%d\n", sumTwos*sumThrees)
}

func countDupsTrips(s string) (int, int) {
	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
	}

	twos := 0
	threes := 0
	for _, v := range m {
		if v == 2 && twos == 0 {
			twos++
		} else if v == 3 && threes == 0 {
			threes++
		}
	}

	return twos, threes
}
