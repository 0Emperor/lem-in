package main

import (
	"fmt"
	"log"
	"os"

	lem "lem/funcs"
)

func main() {
	err, counter := lem.ReadFile(lem.ValidArgs(os.Args))
	if err != "" {
		log.Fatal(err + "1")
	}
	validways := [][]string{}
	chann := make(chan bool)
	go func(s chan bool) {
		os.Stdout.Write(lem.Graphoverview)
		fmt.Println()
		s <- true
	}(chann)
	if counter < 500 {
		validways = lem.Search()
	} else {
		validways = lem.SearchMax()
	}
	<-chann
	if len(validways) == 0 {
		log.Fatal("ERROR: no way found ")
	}
	lem.Sendants(validways)
}
