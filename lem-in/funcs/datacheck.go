package lem

import (
	"log"
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

func ValidData(file *os.File) error {
	data, err := ReadFile(file)
	if err != nil {
		return err
	}
	ended := false
	for i := 0; i < len(data); i++ {
		v := data[i]
		if strings.HasPrefix(v, "#") && !strings.HasPrefix(v, "##") {
			continue
		} else if strings.HasPrefix(v, "##") {
			if v[2:] != "start" && v[2:] != "end" {
				log.Fatal("invalid msg")
			}
		}
		if i == 0 {
			Ants, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal("invalid ants number")
			}
			continue
		}
		if v == "##start" {
			Start = data[i+1]
			Rooms = append(Rooms, strings.Split(data[i+1], " ")[0])
			i++
		} else if v == "##end" {
			End = data[i+1]
			Rooms = append(Rooms, strings.Split(data[i+1], " ")[0])
			ended = true
			i++
		} else if !ended {
			Rooms = append(Rooms, strings.Split(v, " ")[0])
		} else {
			Ways[strings.Split(v, "-")[0]] = append(Ways[strings.Split(v, "-")[0]], strings.Split(v, "-")[1])
			Ways[strings.Split(v, "-")[1]] = append(Ways[strings.Split(v, "-")[1]], strings.Split(v, "-")[0])
		}
	}
	return nil
}
