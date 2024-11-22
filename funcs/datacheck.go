package lem

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ValidArgs(args []string) *os.File {
	if len(args) != 2 {
		return nil
	}
	data, err := os.Open(args[1])
	if err != nil {
		return nil
	}
	return data
}

func ValidData(file *os.File) (error, string) {
	data, err := ReadFile(file)
	if err != nil {
		return err, ""
	}
	counter := 0
	for i := 0; i < len(data); i++ {
		v := data[i]
		if strings.HasPrefix(v, "#") && !strings.HasPrefix(v, "##") {
			continue
		} else if strings.HasPrefix(v, "##") {
			if v[2:] != "start" && v[2:] != "end" {
				log.Fatal("invalid msg")
			}
		}
		Graphoverview = append(Graphoverview, data[i]...)
		Graphoverview = append(Graphoverview, '\n')
		if i == 0 {
			Ants, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal("invalid ants number")
			}
			continue
		}

		if v == "##start" {
			counter++
			if i == len(data)-1 {
				log.Fatal("end or start room missing")
			}
			Start = strings.Split(data[i+1], " ")[0]
			if len(strings.Split(data[i+1], " ")) != 3 {
				log.Fatal("invalid input for start")
			}
			if Start[0] == 'L' {
				log.Fatal("room name cant start with 'L'")
			}
			Graphoverview = append(Graphoverview, data[i+1]...)
			Graphoverview = append(Graphoverview, '\n')
			Rooms = append(Rooms, Start)
			i++
		} else if v == "##end" {
			counter++
			if i == len(data)-1 {
				log.Fatal("end or start room missing")
			}
			End = strings.Split(data[i+1], " ")[0]
			if len(strings.Split(data[i+1], " ")) != 3 {
				log.Fatal("invalid input for end")
			}
			if End[0] == 'L' {
				log.Fatal("room name cant start with 'L'")
			}
			Rooms = append(Rooms, End)
			Graphoverview = append(Graphoverview, data[i+1]...)
			Graphoverview = append(Graphoverview, '\n')
			i++
		} else if strings.Contains(v, " ") {
			counter++
			room := strings.Split(v, " ")[0]
			if len(strings.Split(v, " ")) != 3 {
				log.Fatal("invalid room data.....")
			}
			if room[0] == 'L' {
				log.Fatal("room name cant start with 'L'.....")
			}
			Rooms = append(Rooms, room)
		} else {
			room1 := strings.Split(v, "-")[0]
			room2 := strings.Split(v, "-")[1]
			if slices.Contains(Ways[room1], room2) {
				log.Fatal("2 rooms can only have 1 link between them......")
			}
			counter += 2
			Ways[room1] = append(Ways[room1], room2)
			Ways[room2] = append(Ways[room2], room1)
		}
	}
	if Start == "" || End == "" {
		log.Fatal("end or start room missing....")
	}
	searchmethod := "dfs"
	if counter > 100 {
		searchmethod = "bfs"
	}
	Graphoverview = append(Graphoverview, '\n')
	return nil, searchmethod
}
