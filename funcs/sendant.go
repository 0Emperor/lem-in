package lem

import (
	"fmt"
	"strconv"
)

func ini(s [][]string) map[int]int {
	mapp := make(map[int]int)
	for v := range s {
		mapp[v] = len(s[v])
	}
	return mapp
}

func getmin(v map[int]int) int {
	i := v[0]
	x := 0
	for t, vv := range v {
		if vv < i {
			i = vv
			x = t
		}
	}
	return x
}

func initt(i int) [][]string {
	ff := [][]string{}
	for t := 0; t < i; t++ {
		ff = append(ff, []string{})
	}
	return ff
}

func Sendants(ways [][]string) {
	antgroups := initt(len(ways))
	antid := 1
	sss := ini(ways)

	for antid <= Ants {
		s := getmin(sss)
		antgroups[s] = append(antgroups[s], "L"+strconv.Itoa(antid))
		antid++
		sss[s]++
	}
 controltrafic(antgroups, ways)
}

func controltrafic(antgroups, ways [][]string) {
	trafic := make(map[string]int)
	unavailablerooms := make(map[string]bool)
	finished := []string{}
	for len(finished) != Ants {
		for i := 0; i < len(ways); i++ {
			unavailablerooms[End] = false
			for s := 0; s < len(antgroups[i]); s++ {
				ant := antgroups[i][s]
				if !unavailablerooms[ways[i][trafic[ant]+1]] {
					if ways[i][trafic[ant]+1] == End {
						unavailablerooms[ways[i][trafic[ant]]] = false
						finished = append(finished, ant)
						delete(trafic, ant)
						antgroups[i] = append(antgroups[i][:s], antgroups[i][s+1:]...)
						fmt.Printf("%v-%v ", ant, End)
						s--
						unavailablerooms[End] = true
						continue
					} else {
						fmt.Printf("%v-%v ", ant, ways[i][trafic[ant]+1])
						unavailablerooms[ways[i][trafic[ant]+1]] = true
						unavailablerooms[ways[i][trafic[ant]]] = false
						trafic[ant]++
					}
				}
			}
		}
		fmt.Println()
	}
}