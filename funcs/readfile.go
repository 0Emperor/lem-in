package lem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(file *os.File) (string, int) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	counter := 0
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
				return "ERROR: invalid number of ants (LINE:" + strconv.Itoa(nline) + ")", 0
			}
			if num == 0 {
				return "ERROR: number of ants cant be 0 (LINE:" + strconv.Itoa(nline) + ")", 0
			}
			Ants = num
			i = 1
		} else if i == 2 {
			if !checkroom(line) {
				return "ERROR: invalid start rooom (LINE:" + strconv.Itoa(nline) + ")", 0
			}
			Start = room(line)
			i = 1
		} else if i == 3 {
			if !checkroom(line) {
				return "ERROR: invalid end rooom (LINE:" + strconv.Itoa(nline) + ")", 0
			}
			End = room(line)
			i = 1
		} else if !checkroom(line) {
			if !checklink(line) {
				fmt.Println(line)
				return "ERROR: invalid data format (LINE:" + strconv.Itoa(nline) + ")", 0
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
		Graphoverview = append(Graphoverview, []byte(line)...)
		Graphoverview = append(Graphoverview, '\n')
		counter++
	}
	return "", counter
}

func room(s string) string {
	return strings.Fields(s)[0]
}
