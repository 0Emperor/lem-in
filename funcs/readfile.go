package lem

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(file *os.File) string {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	nline := 0
	for scanner.Scan() {
		nline++
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				i = 2
			} else if line == "##end" {
				i = 3
			}
			continue
		}
		if i == 0 {
			num, err := strconv.Atoi(line)
			if err != nil {
				return "ERROR: invalid number of ants (LINE:" + strconv.Itoa(nline) + ")"
			}
			if num == 0 {
				return "ERROR: number of ants cant be 0 (LINE:" + strconv.Itoa(nline) + ")"
			}
			Ants = num
			i = 1
		} else if i == 2 {
			if !checkroom(line) {
				return "ERROR: invalid start rooom (LINE:" + strconv.Itoa(nline) + ")"
			}
			Start = room(line)
			i = 1
		} else if i == 3 {
			if !checkroom(line) {
				return "ERROR: invalid end rooom (LINE:" + strconv.Itoa(nline) + ")"
			}
			End = room(line)
			i = 1
		} else if !checkroom(line) {
			if !checklink(line) {
				return "ERROR: invalid data format (LINE:" + strconv.Itoa(nline) + ")"
			}
			link := strings.Split(line, "-")
			room1 := link[0]
			room2 := link[1]
			Ways[room1] = append(Ways[room1], room2)
			Ways[room2] = append(Ways[room2], room1)
		} else {
			room := room(line)
			Rooms = append(Rooms, room)
		}
	}
	return ""
}

func room(s string) string {
	return strings.Fields(s)[0]
}
