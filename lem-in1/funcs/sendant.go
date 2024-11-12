package lem

import (
	"fmt"
	"strconv"
	"strings"
)

func Sendants(ways [][]string) {
	antgroups := [][]string{}
	antid := 1
	for i := 0; i < len(ways); i++ {
		antgroup := []string{}
		for j := 0; j < Ants/len(ways); j++ {
			antgroup = append(antgroup, "l"+strconv.Itoa(antid))
			antid++

		}
		if Ants%len(ways) != 0 && i == 0 {
			antgroup = append(antgroup, "l"+strconv.Itoa(antid))
			antid++
		}
		antgroups = append(antgroups, antgroup)
	}
	controltrafic(antgroups, ways)
}

type infos struct {
	steps       int
	currentroom string
}

func controltrafic(antgroups, ways [][]string) {
	trafic := make(map[string]infos)
	unavailablerooms := make(map[string]bool)
	finished := []string{}
	end := strings.Split(End, " ")[0]
	for len(finished) != Ants {
		for i := 0; i < len(ways); i++ {
			for s, ant := range antgroups[i] {
				if trafic[ant].currentroom == end {
					finished = append(finished, ant)
					antgroups[i] = append(antgroups[i][:s], antgroups[i][s+1:]...)
					unavailablerooms[strings.Split(End, " ")[0]] = false
					continue
				}
				if !unavailablerooms[ways[i][trafic[ant].steps+1]] {
					fmt.Printf("%v-%v ", ant, ways[i][trafic[ant].steps+1])
					ii := trafic[ant]
					unavailablerooms[ways[i][trafic[ant].steps+1]] = true
					unavailablerooms[ways[i][trafic[ant].steps]] = false
					ii.currentroom = ways[i][trafic[ant].steps+1]
					ii.steps++
					trafic[ant] = ii
				}
			}
		}
		fmt.Print("\n")
	}
	fmt.Print(finished)
}
