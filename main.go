package main

import (
	"fmt"
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
		validways = lem.SearchMax()
	}
	if len(validways) == 0 || lem.Ants == 0 {
		log.Fatal("ERROR: invalid data format")
	}
	fmt.Println(1111)
	os.Stdout.Write(lem.Graphoverview)
	lem.Sendants(validways)
}
