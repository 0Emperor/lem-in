package lem

import (
	"fmt"
	"strings"
)

func Search() [][]string {
	Dfs(strings.Split(Start, " ")[0])
	sort(solutions)
	fmt.Println(solutions)
	// neededways := gettrials()
	return solutions
}

func sort(unsorted [][]string) [][]string {
	for i := 0; i < len(unsorted); i++ {
		for j := i + 1; j < len(unsorted); j++ {
			if len(unsorted[i]) >= len(unsorted[j]) {
				unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
			}
		}
	}
	return unsorted
}

func gettrials() int {
	fromstart := len(Ways[strings.Split(Start, " ")[0]])
	fromend := len(Ways[strings.Split(End, " ")[0]])
	if fromend <= fromstart {
		return fromend
	} else {
		return fromstart
	}
}

var (
	visited   = make(map[string]bool)
	solutions [][]string
	stack     []string
)

func Dfs(current string) {
	stack = append(stack, current)
	if current == strings.Split(End, " ")[0] {
		solutions = append(solutions, stack)
		 visited[current] = true
	}

	if visited[current] {
		stack = stack[:len(stack)-1]
		return
	}
	visited[current] = true

	for _, v := range Ways[current] {
		Dfs(v)
	}
	stack = stack[:len(stack)-1]
}

// func combine(n int) [][]string {
// }
