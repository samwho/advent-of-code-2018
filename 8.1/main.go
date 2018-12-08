package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	metadata []int
	children []*node
}

type buffer struct {
	pointer int
	data    []int
}

func (b *buffer) next() int {
	n := b.data[b.pointer]
	b.pointer++
	return n
}

func main() {
	fh, _ := os.Open("8.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)
	s.Scan()

	b := &buffer{
		pointer: 0,
		data:    make([]int, 0),
	}

	for _, part := range strings.Split(s.Text(), " ") {
		i, _ := strconv.Atoi(part)
		b.data = append(b.data, i)
	}

	sum := 0
	root := readNode(b)
	traverse(root, func(node *node) {
		for _, metadata := range node.metadata {
			sum += metadata
		}
	})

	fmt.Printf("%d\n", sum)
}

func readNode(b *buffer) *node {
	nchildren := b.next()
	nmetadata := b.next()

	children := make([]*node, 0)
	for i := 0; i < nchildren; i++ {
		children = append(children, readNode(b))
	}

	metadata := make([]int, 0)
	for i := 0; i < nmetadata; i++ {
		metadata = append(metadata, b.next())
	}

	return &node{metadata, children}
}

func traverse(root *node, f func(*node)) {
	f(root)
	for _, child := range root.children {
		traverse(child, f)
	}
}
