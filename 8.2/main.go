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
	fh, _ := os.Open("8.2/data")
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

	root := readNode(b)
	fmt.Printf("%d\n", nodeValue(root))
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

func traverse(root *node, f func(*node) bool) {
	if !f(root) {
		return
	}

	for _, child := range root.children {
		traverse(child, f)
	}
}

func nodeValue(node *node) int {
	sum := 0
	if len(node.children) == 0 {
		for _, metadata := range node.metadata {
			sum += metadata
		}
	} else {
		for _, metadata := range node.metadata {
			index := metadata - 1
			if len(node.children) <= index {
				continue
			}
			sum += nodeValue(node.children[index])
		}
	}
	return sum
}
