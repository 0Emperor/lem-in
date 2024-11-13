package lem

import (
	"fmt"
)

func Search() [][]string {
	Dfs(Start)
	sort(solutions)
	fmt.Println(solutions)
	// neededways := gettrials()
	fmt.Println(conbine(2))
	return conbine(2)
}

func conbine(n int) [][]string {
	
	dd := [][]string{}
	return sort(append(dd, solutions[1], solutions[5]))
}

func collide(p1, p2 []string) bool {
	for i := 1; i < len(p1)-1; i++ {
		for j := 1; j < len(p2)-1; j++ {
			if p2[j] == p1[i] {
				return true
			}
		}
	}
	return false
}

func collides(paths ...[]string) bool {
	for i := 0; i < len(paths)-1; i++ {
		for j := i+1; j < len(paths)-1; j++ {
			if collide(paths[i], paths[j])  {
				return true
			}
		}
	}
	return false
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
	fromstart := len(Ways[Start])
	fromend := len(Ways[End])
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

	if visited[current] { //|| (len(Ways[current]) == 1 && current != End && current != Start)
		stack = stack[:len(stack)-1]
		return
	}

	if current == End {
		supp := []string{}
		supp = append(supp, stack...)
		solutions = append(solutions, supp)
		stack = stack[:len(stack)-1]
		return
	}
	visited[current] = true
	for _, v := range Ways[current] {
		Dfs(v)
	}
	stack = stack[:len(stack)-1]
	visited[current] = false
}

// func combine(n int) [][]string {
// }
