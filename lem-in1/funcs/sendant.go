package lem

import (
	"fmt"
	"strconv"
)

func Sendants(ways [][]string) {
	antgroup := []string{}
	for i := 1; i <= Ants; i++ {
		antgroup = append(antgroup, "l"+strconv.Itoa(i))
	}
	controltrafic(antgroup, ways)
}

type infos struct {
	steps       int
	currentroom string
}

func controltrafic(antgroups []string, ways [][]string) {
	trafic := make(map[string]infos)
	unavailablerooms := make(map[string]bool)
	finished := []string{}
	for len(finished) != Ants {
		for i := 0; i < len(ways); i++ {
			for s, ant := range antgroups {
				if trafic[ant].currentroom == End {
					finished = append(finished, ant)
					antgroups = append(antgroups[:s], antgroups[s+1:]...)
					fmt.Println(111)
				}
				if !unavailablerooms[ways[i][trafic[ant].steps+1]] {

					// fmt.Printf("%v-%v ", ant, ways[i][trafic[ant].steps+1])
					ii := trafic[ant]
					unavailablerooms[ways[i][trafic[ant].steps+1]] = true
					unavailablerooms[ways[i][trafic[ant].steps]] = false
					ii.currentroom = ways[i][trafic[ant].steps+1]
					ii.steps++
					trafic[ant] = ii
				}

			}
		}
		//	fmt.Print("\n")
	}
	fmt.Print(finished)
}
