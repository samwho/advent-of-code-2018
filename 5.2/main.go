package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/vinzmay/go-rope"
)

func main() {
	fh, _ := os.Open("5.2/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	s.Scan()
	str := s.Text()
	min := 50000

	for _, char := range "abcdefghijklmnopqrstuvwxyz" {
		r1 := strings.Replace(str, string(char), "", -1)
		r2 := strings.Replace(r1, string(unicode.ToUpper(char)), "", -1)

		l := len(react(rope.New(r2)).String())

		if l < min {
			min = l
		}
	}

	fmt.Printf("%d\n", min)
}

func react(data *rope.Rope) *rope.Rope {
	i := 0
	for {
		if data.Len() < 2 {
			break
		}

		i++

		if (i + 1) > data.Len() {
			break
		}

		a := data.Index(i)
		b := data.Index(i + 1)

		// fmt.Printf("i: %d, data.Len(): %d, a: %c, b: %c\n", i, data.Len(), a, b)

		if unicode.IsUpper(a) && unicode.IsUpper(b) {
			continue
		}

		if !unicode.IsUpper(a) && !unicode.IsUpper(b) {
			continue
		}

		if unicode.ToLower(a) != unicode.ToLower(b) {
			continue
		}

		data = data.Delete(i, 2)
		i -= 2
		if i < 0 {
			i = 0
		}
	}

	return data
}
