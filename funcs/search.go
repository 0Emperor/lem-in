package lem

import (
	"slices"
)

func Search() [][]string {
	visited[Start] = true
	for _, node := range Ways[Start] {
		Bfs(node, End)
	}
	sort(solutions)
	return solutions
}

func compare(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func sort(slice [][]string) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {

			if len(slice[j]) <= len(slice[i]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
			if compare(slice[j], slice[i]) {
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
}

var (
	visited   = make(map[string]bool)
	solutions [][]string
)

func Bfs(s, e string) {
	queu := []string{s}
	levels := [][]string{Ways[Start]}
	visited[s] = true
	for i := 0; i < len(queu); i++ {
		visiting := queu[i]
		levl := []string{}
		for _, neighbour := range Ways[visiting] {
			if !visited[neighbour] {
				levl = append(levl, neighbour)
				visited[neighbour] = true
				queu = append(queu, neighbour)
			}
			if neighbour == e {
				levels = append(levels, levl)
				solutions = append(solutions, findway(levels))
				return
			}
		}
		levels = append(levels, levl)
	}
}

func contains(d []string, s [][]string) bool {
	for _, v := range s {
		for _, f := range d {
			if !slices.Contains(v, f) {
				return false
			}
		}
	}
	return true
}

func findway(levels [][]string) []string {
	curent := End
	visited = make(map[string]bool)
	way := []string{curent}

	for i := len(levels) - 1; i >= 0; i-- {
		for _, v := range levels[i] {
			if exist(curent, Start) && !contains(way, solutions) {
				way = append(way, Start)
				return flip(way)
			}
			if exist(curent, v) {
				way = append(way, v)
				visited[curent] = true
				curent = v
			}
			if curent == Start {
				return flip(way)
			}
		}
	}
	way = append(way, Start)
	return flip(way)
}

func flip(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func exist(s, v string) bool {
	for _, t := range Ways[s] {
		if t == v {
			return true
		}
	}
	return false
}
