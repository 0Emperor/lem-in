package lem

import (
	"fmt"
	"slices"
)

func SearchMax() [][]string {
	visited[Start] = true
	for i := len(Ways[Start]) - 1; i >= 0; i-- {
		node := Ways[Start][i]
		Bfs(node, End)
	}
	sort(solutions)
	validate(solutions)
	return solutions
}

func validate(s [][]string) {
	for i := 0; i < len(s); i++ {
		e := s[i][1 : len(s[i])-1]
		for j := i + 1; j < len(s); j++ {
			r := s[j][1 : len(s[j])-1]
			if compare2(r, e) {
				if i == len(s)-1 {
					s = s[:i]
				} else {
					s = append(s[:j], s[j+1:]...)
					fmt.Println(s)
				}
			}
		}
	}
	solutions = s
}

func compare2(s1, s2 []string) bool {
	ss := s1
	s := s2
	if len(s1) < len(s2) {
		ss = s2
		s = s1
	}
	for i, v := range s {
		if ss[i] == v {
			return true
		}
	}
	return false
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

func sort(slice [][]string) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[j]) <= len(slice[i]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
			if compare(slice[j][1:len(slice[i])-1], slice[i][1:len(slice[i])-1]) {
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
}
func Bfs(s, e string) {
	fmt.Println(s, visited)
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

			if exist(curent, Start) && !contains(append(way, Start), solutions) {
				way = append(way, Start)
				visited[curent] = true
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
