package lem

import (
	"os"
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

/*
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

				if !checkroom(data[i+1]) {
					log.Fatal("Invalid room3")
				}
				Start = strings.Split(data[i+1], " ")[0]
				Graphoverview = append(Graphoverview, data[i+1]...)
				Graphoverview = append(Graphoverview, '\n')
				Rooms = append(Rooms, Start)
				i++
			} else if v == "##end" {
				counter++
				if i == len(data)-1 {
					log.Fatal("end or start room missing")
				}

				if !checkroom(data[i+1]) {
					log.Fatal("Invalid room2")
				}
				End = strings.Split(data[i+1], " ")[0]
				Rooms = append(Rooms, End)
				Graphoverview = append(Graphoverview, data[i+1]...)
				Graphoverview = append(Graphoverview, '\n')
				i++
			} else if strings.Contains(v, " ") {
				counter++
				if !checkroom(v) {
					log.Fatal("Invalid room1")
				}
				room := strings.Split(v, " ")[0]

				Rooms = append(Rooms, room)
			} else {
				link := strings.Split(v, "-")
				if len(link) != 2 {
					log.Fatal("invalid link")
				}
				room1 := link[0]
				room2 := link[1]
				if !seen[room1] || !seen[room2] {
					log.Fatal("link contains inexistentroom")
				}
				if slices.Contains(Ways[room1], room2) {
					log.Fatal("2 rooms can only have 1 link between them......")
				}
				counter += 1
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
*/
func checklink(link string) bool {
	linnk := strings.Split(link, "-")
	if len(linnk) != 2 {
		return false
	}
	if !seen[linnk[1]] || !seen[linnk[0]] {
		return false
	}
	return true
}

func checkroom(room string) bool {
	roomparts := strings.Fields(room)
	if len(roomparts) != 3 {
		return false
	}
	_, err1 := strconv.Atoi(roomparts[1])
	if _, err2 := strconv.Atoi(roomparts[2]); err1 != nil || err2 != nil {
		return false
	}
	roomname := roomparts[0]
	if roomname[0] == 'L' {
		return false
	}
	if seen[roomname] {
		return false
	}
	seen[roomname] = true
	return true
}
