package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	m := make(map[int]bool)

	for {
		fh, err := os.Open("2/2.data")
		if err != nil {
			panic(err)
		}
		defer fh.Close()
		s := bufio.NewScanner(fh)

		for s.Scan() {
			i, err := strconv.Atoi(s.Text())
			if err != nil {
				panic(err)
			}

			sum += i
			_, ok := m[sum]
			if ok {
				fmt.Printf("%d\n", sum)
				return
			}

			m[sum] = true
		}
	}
}
