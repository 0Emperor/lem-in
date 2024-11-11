package lem

import (
	"strings"
)

func Search() [][]string {
	validways := [][]string{}
	Trials := gettrials()
	for i := 0; i < Trials; i++ {
		way, deadend := Bfs()
		if !deadend {
			validways = append(validways, way)
		}
	}
	return validways
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

func Bfs() ([]string, bool) {
	levels := [][]string{}
	visited := make(map[string]bool)
	visited[strings.Split(Start, " ")[0]] = true
	tovisit := []string{strings.Split(Start, " ")[0]}
	for i := 0; i < len(tovisit); i++ {
		levl := []string{}
		visiting := tovisit[i]
		visited[visiting] = true
		for _, v := range Ways[visiting] {
			if !visited[v] {
				tovisit = append(tovisit, v)
				visited[v] = true
				levl = append(levl, v)
			}
			if v == strings.Split(End, " ")[0] {
				goto out
			}

		}
		levels = append(levels, levl)
	}
	return nil, true
out:

	return findway(levels), false
}

func findway(levels [][]string) []string {
	curent := strings.Split(End, " ")[0]
	way := []string{curent}
	for i := len(levels) - 1; i >= 0; i-- {
		for _, v := range levels[i] {
			if exist(curent, v) {
				way = append(way, v)
				curent = v
			}
		}
	}
	closeways(way[1:])
	way = append(way, strings.Split(Start, " ")[0])

	return flip(way)
}
func flip(s []string)[]string  {
	r:=[]string{}
	for i := len(s)-1; i >= 0; i-- {
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

func closeways(way []string) {
	for _, room := range way {
		for _, v := range Ways[room] {
			i := -1
			for o, x := range Ways[v] {
				if x == room {
					i = o
					break
				}
			}
			if i != -1 {
				Ways[v] = append(Ways[v][:i], Ways[v][i+1:]...)
			}
		}
	}
}
