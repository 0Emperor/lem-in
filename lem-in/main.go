package main

import (
	"os"

	lem "lem/funcs"
)

func main() {
	lem.ValidData(lem.ValidArgs(os.Args))
	validways := lem.Search()
	lem.Sendants(validways)
}
