package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fh, _ := os.Open("1.data")
	defer fh.Close()
	s := bufio.NewScanner(fh)
	sum := 0

	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		sum += i
	}

	fmt.Printf("%d\n", sum)
}
