package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fh, _ := os.Open("4/4.data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	ss := make([]string, 0, 250)
	for s.Scan() {
		ss = append(ss, s.Text())
	}

	for _, a := range ss {
		for _, b := range ss {
			c := commonChars(a, b)
			lendiff := len(a) - len(c)

			if lendiff == 1 {
				fmt.Printf("%s\n", c)
			}
		}
	}
}

func commonChars(a, b string) string {
	var builder strings.Builder
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			builder.WriteByte(a[i])
		}
	}
	return builder.String()
}
