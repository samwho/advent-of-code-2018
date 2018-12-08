package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fh, _ := os.Open("7.1/data")
	defer fh.Close()
	s := bufio.NewScanner(fh)

	m := make(map[string]map[string]bool)

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")

		prereq := parts[1]
		task := parts[7]

		_, ok := m[task]
		if !ok {
			m[task] = make(map[string]bool)
		}
		_, ok = m[prereq]
		if !ok {
			m[prereq] = make(map[string]bool)
		}

		m[task][prereq] = true
	}

	var order strings.Builder
	for {
		tasks := canBeDone(m)
		if len(tasks) == 0 {
			break
		}
		sort.Strings(tasks)
		do(tasks[0], m)
		order.WriteString(tasks[0])
	}

	fmt.Printf("%s\n", order.String())
}

func canBeDone(m map[string]map[string]bool) []string {
	ret := make([]string, 0)
	for task, prereqs := range m {
		if len(prereqs) == 0 {
			ret = append(ret, task)
		}
	}
	return ret
}

func do(task string, m map[string]map[string]bool) {
	for _, prereqs := range m {
		delete(prereqs, task)
	}
	delete(m, task)
}
