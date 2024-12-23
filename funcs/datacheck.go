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

func checklink(link string) bool {
	linnk := strings.Split(link, "-")
	if len(linnk) != 2 {
		return false
	}
	if checklinks[link] || checklinks[linnk[1]+"-"+linnk[0]] {
		return false
	}
	if !seen[linnk[1]] || !seen[linnk[0]] {
		return false
	}
	checklinks[link] = true
	checklinks[linnk[1]+"-"+linnk[0]] = true
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
