package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type worker struct {
	task          string
	timeRemaining int
}

func main() {
	fh, _ := os.Open("7.2/data")
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

	workers := make([]worker, 5, 5)
	var totalTime int

	for {
		tasks := canBeDone(m, workers)
		sort.Strings(tasks)

		for _, task := range tasks {
			worker := findWorker(workers)
			if worker == -1 {
				break
			}

			workers[worker].timeRemaining = 60 + int(task[0]) - int('A') + 1
			workers[worker].task = task
		}

		for i, worker := range workers {
			if worker.timeRemaining > 0 {
				workers[i].timeRemaining--
				if workers[i].timeRemaining == 0 {
					do(workers[i].task, m)
				}
			}
		}

		totalTime++

		if len(m) == 0 {
			break
		}
	}

	fmt.Printf("%d\n", totalTime)
}

func canBeDone(m map[string]map[string]bool, workers []worker) []string {
	ret := make([]string, 0)
	for task, prereqs := range m {
		if len(prereqs) == 0 && !isInProgress(task, workers) {
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

func findWorker(workers []worker) int {
	for i, worker := range workers {
		if worker.timeRemaining == 0 {
			return i
		}
	}

	return -1
}

func isInProgress(task string, workers []worker) bool {
	for _, worker := range workers {
		if worker.task == task {
			return true
		}
	}
	return false
}
