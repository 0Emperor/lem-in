package lem

import (
	"fmt"
)

func SearchMax() [][]string {
	visited[Start] = true
	for i := len(Ways[Start]) - 1; i >= 0; i-- {
		node := Ways[Start][i]
		Bfs(node, End)
	}
	sort1(solutions)
	return solutions
}

func compare(s1, s2 []string) bool {
	if len(s1) != len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func sort1(slice [][]string) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[j]) <= len(slice[i]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
			if compare(slice[j][1:len(slice[i])-1], slice[i][1:len(slice[i])-1]) {
				slice = append(slice[:i], slice[i+1:]...)
			}
			if !hhhh(slice[j][1:len(slice[i])-1], slice[i][1:len(slice[i])-1]) {
				if len(slice[j]) < len(slice[i]) {
					slice = append(slice[:i], slice[i+1:]...)
				} else {
					slice = append(slice[:j], slice[j+1:]...)
				}
			}
		}
	}
	solutions = slice
}

func Bfs(s, e string) {
	fmt.Println(s, visited)
	parent := make(map[string]string)
	parent[s] = Start
	queu := []string{s}
	visited[s] = true
	for i := 0; i < len(queu); i++ {
		visiting := queu[i]
		for _, neighbour := range Ways[visiting] {
			if !visited[neighbour] {
				visited[neighbour] = true
				parent[neighbour] = visiting
				queu = append(queu, neighbour)
			}
			if neighbour == e {
				solutions = append(solutions, findway(parent))
				return
			}
		}
	}
}

func findway(parent map[string]string) []string {
	curent := End
	visited = make(map[string]bool)
	way := []string{curent}
	for curent != Start {
		way = append(way, parent[curent])
		curent = parent[curent]
		visited[curent] = true
	}
	return flip(way)
}

func flip(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
