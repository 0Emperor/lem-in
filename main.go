package main

import (
	"log"
	"os"

	lem "lem/funcs"
)

func main() {
	err, Search := lem.ValidData(lem.ValidArgs(os.Args))
	if err != nil {
		log.Fatal(err)
	}
	validways := [][]string{}
	switch Search {
	case "bfs":
		validways = lem.SearchMax()
	default:
		validways = lem.Search()
	}
	if len(validways) == 0 || lem.Ants == 0 {
		log.Fatal("ERROR: invalid data format")
	}
	os.Stdout.Write(lem.Graphoverview)
	lem.Sendants(validways)
}
