package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/vinzmay/go-rope"
)

func main() {
	fh, _ := os.Open("5.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	s.Scan()
	data := rope.New(s.Text())

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

	fmt.Printf("%s\n", data.String())
	fmt.Printf("%d\n", len(data.String()))
}
